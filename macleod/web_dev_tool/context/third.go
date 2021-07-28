package main

import (
        "fmt"
        "log"
        "net/http"
        "time"
        "context"
)

func main() {
     http.HandleFunc("/", foo)
     http.HandleFunc("/bar", bar)
     http.Handle("/favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
     ctx := req.Context()
     ctx = context.WithValue(ctx, "uid", 9999) 
     ctx = context.WithValue(ctx, "fName", "Bond") 
   
     results, err := dbAccess(ctx)
     if err != nil {
        http.Error(w , err.Error(), http.StatusRequestTimeout)
     }
     fmt.Fprint(w, results)
    
}

func dbAccess(ctx context.Context) (int, error) {
    ctx, cancel :=  context.WithTimeout(ctx, 3 * time.Second)
    defer cancel()
    
    ch := make(chan int)
    go func() {
       num := ctx.Value("uid").(int)
       time.Sleep(10 * time.Second)
   
       if ctx.Err() != nil {
           return
       }     
  
      ch <- num
    
    }()

    select {
      case <-ctx.Done():
       return 0, ctx.Err()
      case i := <-ch:
          return i, nil
    } 
    
}


func bar(w http.ResponseWriter, req *http.Request) {
     ctx := req.Context()
     log.Println(ctx)
     fmt.Fprint(w, ctx)
}

package main

import (
        "fmt"
        "net/http"
        "log"
        "context"
)

func main() {
     
     http.HandleFunc("/", foo)
     http.HandleFunc("/bar", bar)
     http.Handle("/favicon.ico", http. NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
     ctx := req.Context() 
     ctx =  context.WithValue(ctx, "userId" , 999) 
     ctx =  context.WithValue(ctx, "fName" , "James")
     
     results := dbAccess(ctx) 
     fmt.Fprintln(w, results)
}

func dbAccess( ctx context.Context) int {
     num := ctx.Value("userId").(int)
     return num
}

func bar(w http.ResponseWriter, req *http.Request) {
     ctx := req.Context()
     log.Println(ctx)
     fmt.Fprintln(w, ctx)
}

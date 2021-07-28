package main

import (
         "net/http"
         "advance"
         "golang.org/x/net/context"
)

func main() {
     logger := advance.CreateLog("../webserver") 
     http.Handle("/", advance.Next(logger, hello))
     http.Handle("/panic", advance.Recover(rec))
     http.Handle("/context", advance.PassContext(withContext))
     http.ListenAndServe(":3333", nil)

}

func hello(w http.ResponseWriter, req *http.Request) {
     w.Write([]byte("Hello GO World!"))
}

func rec(w http.ResponseWriter, req *http.Request) {
     panic(advance.ErrInvalidID)
}

func withContext(ctx context.Context, w http.ResponseWriter, req *http.Request) {
    val :=  ctx.Value("foo") 
    w.Write(( val.([]byte)))
}

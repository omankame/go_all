package main

import (
        "fmt"
        "middleware"
        "net/http"
)

func main() {
     sum := middleware.Add(1, 2, 3)
     fmt.Println(sum)


// & as middleware is other package
     chain := &middleware.Chain{
                        0 }    

    sum2 := chain.AddNext(1).AddNext(2).AddNext(3).Finally(0)
    fmt.Println(sum2) 

    http.Handle("/", middleware.NEXT(hello))
    http.ListenAndServe(":3333", nil)
} 

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing...")
        w.Write([]byte("HELLO"))
}

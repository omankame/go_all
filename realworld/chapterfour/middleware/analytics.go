package middleware

import ( 
        "fmt"
         "net/http"
)

//variadic func

func Add(num ...int) int {
     sum := 0
     for _, n := range num {
         sum += n
     }
    return sum

}

// structure that  hold value

type Chain struct {
     Sum int

}

func (c *Chain) AddNext(num int) *Chain {
     c.Sum += num
     return c
}

func (c *Chain) Finally(num int) int {
     return c.Sum + num
}

func NEXT(next http.HandlerFunc) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
           fmt.Println("Before")
           next.ServeHTTP(w,r)
           fmt.Println("After")
           })
}

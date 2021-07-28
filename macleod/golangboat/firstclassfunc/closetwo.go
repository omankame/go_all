package main

import (
         "fmt"
)

func appendString() func(string) string {
     t := "Hello"
     c := func(s string) string {
         t = t + " " + s
         return t
      }

     return c

}
   
func main() {

     a := appendString()
     b := appendString()

     fmt.Println(a("World"))
     fmt.Println(b("Everyone"))

     fmt.Println(a("Gopher"))
     fmt.Println(b("!"))
}


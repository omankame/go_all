package main

import (
         "fmt"
)

func main() {
     
     a := func() {
               fmt.Println("Hello, This is first example of Anonymous Function")
          }

     a()

     fmt.Printf("%T\n", a)

}

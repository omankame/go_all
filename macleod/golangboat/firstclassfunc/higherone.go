package main

import (
        "fmt"
)

func main() {

     f := func(a int, b int) int {
          return a + b
         }

     simple(f)

}
func simple(z func(a, b int) int) {
       x := z(90, 50)
       fmt.Printf("Sum is %[1]d and type is %[1]T\n", x)
} 
         
         

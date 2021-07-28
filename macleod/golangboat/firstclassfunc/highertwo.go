package main

import (
        "fmt"
)

func simple() func(a int, b int) int {
     f := func(a int, b int) int {
          return a + b
     }

    return f
}

func main() {
     s := simple()

     fmt.Println(s(90, 50))
}



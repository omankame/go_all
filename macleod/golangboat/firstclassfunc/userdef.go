package main

import (
        "fmt"
)

type add func(a int, b int) int

func main() {
     
     var z add = func(a int, b int) int {
                 return a + b
         }

    s :=  z(50, 90)
    fmt.Println("Sum is ", s)
}

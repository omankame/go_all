package main

import (
         "fmt"
)

func main() {
    outerloop:
     for i := 0; i < 3 ; i++ {
        for j := 1; j < 5 ; j++ {
           fmt.Printf("i = %d, j = %d\n", i,j)
           if i == j {
            continue outerloop
           }
         }
     }
}
           

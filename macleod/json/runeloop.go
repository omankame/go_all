package main

import (
          "fmt"
)

func main() {
     c := 'A'
     asciiValue := int(c)
     var a string
     for i:= 0; i < 5; i++ {
       a = string(asciiValue) 
       fmt.Println(a + "1")
       asciiValue++
     }


    rlast := 1
    c  = 'O'
    asciiValue = int(c)
    var b string
    for i := 0; i < 5; i++ {
     b = string(asciiValue)
     concat :=  fmt.Sprint(b, rlast)
     fmt.Println(concat)
     asciiValue++
    }
}     

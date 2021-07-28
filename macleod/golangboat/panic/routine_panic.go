package main

import (
         "fmt"
)

func main() {

     a := 5 
     b := 0
     cal(a,b)

    fmt.Println("Finally closed from main routine")
}

func cal(a int, b int) {
     defer recovery()
     fmt.Println( a + b)
     chn := make(chan bool)
     divide(a, b, chn)
     <-chn
}

func recovery() {
    if r := recover(); r != nil {
       fmt.Println("Recovered:", r)
    }
}

func divide(a int, b int, chn chan bool) {
     res := a/b
     fmt.Println(res)
     chn <- true
}    

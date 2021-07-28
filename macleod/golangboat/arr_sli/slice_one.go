package main

import (
       "fmt"
)

func subFunc(a []int) []int {
     for i := range a {
     a[i] -= 2
     }
     return a
}

func main() {
     var a = []int{10, 12, 15}
     fmt.Println("SLice before function call", a)
     b := subFunc(a)
     fmt.Println("SLice after function call", b)
}

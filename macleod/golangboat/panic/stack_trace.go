package main

import (
        "fmt"
        "runtime/debug"
)

func invalidAccess() {
     defer recoverSlice()
   
     n := []int{3, 4, 5}
     fmt.Println(n[4])
}

func main() {
     invalidAccess()
     fmt.Println("Finally closed from main sucessfully")
}

func recoverSlice() {
     if r := recover(); r != nil {
        fmt.Println("Recovered", r)
      }
    debug.PrintStack()
}     

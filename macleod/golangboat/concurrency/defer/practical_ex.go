package main

import (
       "fmt"
       "sync"
)

type rect struct {
     length   int
     width    int
}

func (r rect) area( wg *sync.WaitGroup) {
   
     defer wg.Done()

     if r.length < 0 {
        fmt.Println("The length of the rect should be positive")
        return
     }

     if r.width < 0 {
        fmt.Println("The length of the width should be positive")
        return
     }

     area := r.length * r.width
     fmt.Println("The area of the rectangle is :", area)     
}

func main() {

     var wg sync.WaitGroup
     re1 := rect{-20, 20,}
     re2 := rect{20, 20,}
     re3 := rect{5, 20,}
     re4 := rect{20, -20,}
     re5 := rect{40, 20,}

     re := []rect{re1, re2, re3, re4, re5 }

     for _, v := range re {
         wg.Add(1)
         go v.area(&wg)
      }

      wg.Wait()
      fmt.Println("SO final calculation done with help of goroutine for all 5 instances") 
     
}

package main

import (
        "fmt"
)

func main() {

     sl := []int{5,6,7,8,9,10}

     fr := iMap(sl, func(n int) int {
              return n * n
       })

    fmt.Println(fr)
}

func iMap(s []int, f func(int) int)  []int {

     var r []int

     for _, v := range s {
           
        r = append(r, f(v))
     }

     return r
}
      

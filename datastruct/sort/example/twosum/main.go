package main

import (

       "fmt"
       "sort"
)

func main() {
     a := []int{1, 4, 45, 6, 10, -8}   // sum should be 16

     sort.Ints(a)
     fmt.Println(a)

     //method 1 l and r pointer 

     l,r  := 0, len(a) -1

     i := 0
     for i < len(a) -1 {
         if a[l] + a[r] > 49 {
            r--
         } else if a[l] + a[r] < 49 {
           l++
         } else {
           fmt.Println(a[l], a[r], a[l] + a[r])
           break
         }
        
         i++
     }
     

}     


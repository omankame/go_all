package main

import (
           "fmt"
)

func main() {
     a := []int{1, 33, 4, 1,   8, 99, 2}
     num := firstrecurrenceoccur(a)   
     if num == 0 {
        fmt.Println("No repeatative number in array")
     } else {
       fmt.Println("Reoccurence of number is:", num)
     }
}

func firstrecurrenceoccur(a []int) int {
     m := make(map[int]bool)

     for _, v := range a {
         
        if _,  ok := m[v]; ok {
           return v 
        }

         m[v] = true
     }


     return 0
     
}

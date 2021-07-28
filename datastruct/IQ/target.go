package main

import (
        "fmt"
        
)

func main() {
     a := []int{2, 6, 3, 8, 4} //target 10

     for i, _ := range a {
         for j :=i+1; j < len(a); j++ {
             if a[i] + a[j] == 11 {
               fmt.Println("Target gain ", a[i]," and", a[j])
               return
             }  
         }

     }
} 
    


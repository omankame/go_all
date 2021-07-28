package main

import (
        "fmt"
)

func main() {

     a := []int{44, 12, 2, 456, 99, 66, 3, 21}

     fmt.Println(bubbleSort(a, len(a)) )
     
}

func bubbleSort(a []int, n int) []int {
     if n == 1 {
        return a
     }

     var i = 0

     for i < n-1 {
         if a[i] > a[i+1] {
           a[i], a[i+1] = a[i+1], a[i]
         }
         i++
     }

     bubbleSort(a, n -1) //recursive call to same function 
     return a
}   
         



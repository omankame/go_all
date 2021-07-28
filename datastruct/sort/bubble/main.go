package main

import (
        "fmt"
)

func main() {
    arr := []int{9, 3, 22, 1, 99, 4, 13}
    
   s :=  sorting(arr)
   fmt.Println(s)

}

func sorting(arr []int) []int {
     
     n := len(arr) -1 // bubble sort number of passes should be n -1, where n is length of arra

     for i := 0; i < n - 1 ; i++ { // outer loop for passes
         for j := 0; j < n - 1; j++ {  // inner loop for comparison
             if arr[j] > arr[j+1] {
                  temp := arr[j]
                  arr[j] = arr[j+1]
                  arr[j+1] = temp
             }
         }
     }
 
      return arr
} 



package main

import (
        "fmt"
)

func main() {
    arr := []int{9, 3, 22, 111, 999, 4444, 133333}
    
   s :=  sorting(arr)
   fmt.Println(s)

}

func sorting(arr []int) []int {
     
     n := len(arr) -1 // bubble sort number of passes should be n -1, where n is length of arra

     for i := 0; i < n - 1 ; i++ { // outer loop for passes
              flag := true
        for j := 0; j < n - 1 -i; j++ {  // inner loop for comparison and calclulation reduce as last elemnts are sorted and no action required
             if arr[j] > arr[j+1] {
                  temp := arr[j]
                  arr[j] = arr[j+1]
                  arr[j+1] = temp
                  flag = false   // if swapping happen in pass then flag become false
             }
         }
     
           if flag == true {  // if swapping not done thoughout the entire pass it means array is sorted and no further action required
              return arr
              break
           } 

      }
 
      return arr
} 



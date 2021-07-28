package main

import (
         "fmt"
)

func main() {
     a := []int{22, 7, 8, 33, 90, 5, 2, 12}
     n := len(a)
     for i := 0; i < n-1; i++ {  // i strt from 0th position, and n is number of passes len(n) - 1
         min := i  // min element use for swapping nad pointer fro element comparison
         for j := i+1; j < n;  j++ { // j strt from next element hence j =1 and iteration is less than n for inner loop
             if a[j] < a[min] {  // compare value and assign index to min if condition satisfy
                min = j
             }
         }

       if min != i {  //it means value found in unsorted array which is less than i index value
          a[i], a[min] = a[min], a[i]
       }

     }

    fmt.Println(a)
}

         

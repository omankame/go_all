package main

import (
           "fmt"
)

func main() {
     a := []int{1, 33, 3, 1,  8, 99, 2}
     num := firstrecurringCharacter(a) 
     if num == 0 {
          fmt.Println("No Recurring Character")
          return
     }
 
     fmt.Println("First Recurring Character", num)
}

func firstrecurringCharacter( a []int) int {

     for i := 0; i < len(a); i++ {
         for j := i+1 ; j < len(a); j++ {
             if a[i] == a[j]  {
                return a[i]
                break
             } else if a[j] == a[j+1] {
                    return a[j]
                    break
             }      
          }
      }
  
      return 0
}


// soln time complexity is o(n^2) if map use then it get reduce. Hence second solution

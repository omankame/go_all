package main 

import (
         "fmt"
)

func main() {
     a := []int{3, 44, 4, 6, 77, 11, 1}
     fmt.Println(insertSort(a))
}

func insertSort(a []int) []int {
     

     for i := 1; i < len(a)   ; i++ {
            tmp := a[i]
            fmt.Println("loop", i)
        for  j := i - 1; j >= 0; j-- {
             fmt.Println(j)
           if tmp <  a[j] {
            a[j+1] = a[j]
            } else {
             a[j+1] = tmp
             break
            }
                    
        }
          
      
     }
   return a
}
         
     

package main

import (
        "fmt"
)

func main() {
    a := []int{1, 2, 1, 39, 2, 9, 7, 54, 11}
 
    fmt.Println(countSort(a))

}

func countSort(a []int) []int {
     n := len(a)
     output := make([]int, n)  // to keep the final ouput


     // to calculate the K range value

    max := a[0]
    i := 1
    for i < n {
        if a[i] > max {
           max = a[i]
        }
        i++
    }

    count := make([]int, max + 1)  // middle count array declaration


    j := 0
    for j < n {
       count[ a[j] ]++    //count array values logic
    
       j++
    }
    fmt.Println(count)


    // count aray valur update logic 

    k := 1

    for k < len(count) {
        count[k] = count[k] + count[k-1]  // seein notebook
        k++
    }

    fmt.Println(count)

    m := 0
    for m < n {
        output[ count[ a[m] ] -1 ] = a[m]
        count[a[m]]--
        m++
    }
   
    return output

}
       

    
    
     


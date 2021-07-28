package main

import (
         "fmt"
         "sort"
)

func main() {

     a := [] int{0, 3, 4, 31}
     b := [] int{4, 6, 30}

     mergeSort(a, b)
}

func mergeSort(a []int, b []int) {

      a = append(a, b...)  //byusing built in function in go is simplest method to sort
      sort.Ints(a)
      fmt.Println(a) 
}
     


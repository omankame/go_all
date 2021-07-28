package main

import (
        "fmt"
)

func main() {

     a := []string{"A", "B", "C", "D", "E"}
     i := len(a) //number of index insert X

//     a = a[:i] // A B 
//     a = a[i:] // C D E
//     a = append([]string{"X"}, a[i:]...) // X C D E
     
       a = append(a[:i], append([]string{"X"}, a[i:]...)...) 

     fmt.Println(a)

}

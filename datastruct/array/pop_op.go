package main

import (
        "fmt"
)

func main() {

     a := []string{"A", "B", "C", "D", "E"}
     i := 2

     copy(a[i:], a[i+1:]) // op [A B D E E]
     a[len(a) -1] = ""    // op [A B D E  ]
     
     fmt.Println(a)
     a = a[:len(a) -1]
     fmt.Println(a)  //final op [A B D E]

}

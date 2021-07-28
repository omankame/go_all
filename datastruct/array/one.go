package main

import (
        "fmt"
)

//https://golang.org/src/go/types/type.go?s=1709:1753#L84



type Data struct {
     index int
     val   int
}

type Array struct {
     len  int
     data Data 
} 

func NewValue() *Array {
     return &Array {
               len: 0,
               data:  Data {
                     index: -1,
                     val:    0,
               },
      }
}

var da Array = *NewValue()
 

func main() {
    ar := Array{}
      ar.push(5)        
    ar.push(6)
}         

func (a Array) push (n int)  {
      a.data.index++
      a.data.val = n
   
      a.len++

      fmt.Printf("Length : %d, data: { %d: %d} \n", a.len, a.data.index, a.data.val)
      fmt.Println(a)
} 
 

                    
     


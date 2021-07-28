package main

import (
        "fmt"
)

//https://golang.org/src/go/types/type.go?s=1709:1753#L84

type Type interface {
	Underlying() Type

//	String() string
}


type Array struct {
	len  int
	elem int
} 

func (a Array) Underlying() Type {
     return a
}




func main() {
 
    ar := Array{}
    var t Type
    fmt.Println(ar)
    t = ar
    fmt.Println(t.Underlying()) 

     ar.Push(5)    
}         
 
func (a Array) Push(n int)  {
     a.elem = n
     a.len++
     fmt.Println(a)
}

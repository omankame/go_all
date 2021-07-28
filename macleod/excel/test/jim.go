package main

import (
      "fmt"
)

func main() {
//    s:=  fmt.Sprintf("Hello World\nHow are you doing today\nHope all is well with your go\nAnd code")

//    fmt.Println(s)

    st := []string{"om", "uttu", "bittu"}

    var final string
    for _, v := range st {
   
    final += v 
    final = final + "\n"
    }

    fmt.Println(final)
    
}

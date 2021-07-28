package main

import (
        "fmt"
)

func main() {

     name := "onkar"

    final :=  rotate(name)
    fmt.Println(string(final))


}

func rotate(txt string) []byte {

     b := make([]byte,0, len(txt))
     c := []byte(txt)
     for _, v := range c {

     if v < 110 {
        v = v + 13

      } else {
        v = v - 13
      }
     b = append(b, v)

   }

   return b  


}

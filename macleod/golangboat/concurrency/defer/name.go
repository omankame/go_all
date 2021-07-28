package main

import (
        "fmt"
)

func main() {
     name := "Onkar"

     fmt.Printf("The origina name: %s\n", name)
     defer fmt.Println()

     for _, v := range []rune(name) {
         defer fmt.Printf("%c", v)
     }

      
}




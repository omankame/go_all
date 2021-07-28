package main

import (
       "fmt"
)

func main() {
     
    switch n := num(); {
    case n < 50:
          fmt.Println("The number is less than 50")
    fallthrough
    case n < 100:
          fmt.Println("The number is less than 100")
    fallthrough
    case n < 150:
          fmt.Println("The number is less than 150")
 
   }

}

func num() int {
     return 15 * 5 
}
    

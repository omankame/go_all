package main

import (
         "fmt"
)

func main() {
   
     str := "I am Onkar"
     fmt.Println(str)
     revStr(str)
     fmt.Println()
}

/* func revStr(str string) {

     for _, v := range str {
         defer fmt.Printf("%[1]v", string(v))
     }
} */

func revStr(str string) {

     revArray := []rune{}
     lenStr := len(str) - 1 
     for i := lenStr; i >= 0; i-- {
         revArray = append(revArray, rune(str[i]))
     }

     fmt.Print(string(revArray))

}



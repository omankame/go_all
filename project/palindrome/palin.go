package main

import (
        "fmt"
        "strings"
        "bufio"
        "os"
)

func main() {

     reader := bufio.NewReader(os.Stdin)
     fmt.Println("Type the string")

     p_string, err := reader.ReadString('\n')
     if err != nil {
     fmt.Println(err)
     os.Exit(1)
     }
     p_string = strings.TrimSuffix(p_string, "\n")
     fmt.Printf("The string is %s \n", p_string) 

     r_string := reverse(p_string)
     
     fmt.Printf("The reverse string is %s \n", r_string)

     palin_rs := p_string == r_string 

     if palin_rs {
        fmt.Println("Given string is palindrome")
     } else {
        fmt.Println("Given string is not palindrome")
     }
}

func reverse(p string) (res string) {
     for _, v := range p {
         res = string(v) + res
      }
     return
}
  


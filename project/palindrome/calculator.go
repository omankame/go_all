package main

import (
        "fmt"
        "os"
        "strconv"
        "bufio"
        "strings"
)

func main() {

     args := os.Args[1:]
     if len(args) != 2 {
     fmt.Println("Please enter two digits for Calculation")
     return
     }

     num1, err := strconv.Atoi(args[0])
     if err != nil  || num1 <= 0 {
     fmt.Println("Please enter only valid digit.")
     return
     }
     
     num2, err := strconv.Atoi(args[1])
     if err != nil  || num2 <= 0 {
     fmt.Println("Please enter only valid digit.")
     return
     }
     for {
     reader := bufio.NewReader(os.Stdin)
     fmt.Printf("Type Command for desired Operation.  1. ADD 2. SUB 3. DIV  4. MUL  5. EXIT \n")
     str, err := reader.ReadString('\n') 
     if err != nil {
     fmt.Println("Please enter valid Operations Only")
     continue
     }
     str = strings.TrimSuffix(str, "\n") 
    
     switch str {
      
     case "ADD":
           res := num1 + num2
           fmt.Printf("The addition of two number is: %d \n", res)

     case "SUB":
           
           res := num1 - num2
           fmt.Printf("The substraction of two number is: %d \n", res)
      
     case "DIV":
           res := float64(num1) / float64(num2)
           fmt.Printf("The division  of two number is: %.2f \n", res)

     
     case "MUL":
           res := num1 * num2
           fmt.Printf("The multipication of two number is: %d \n", res)
     
     case "EXIT":
           return
       

     } 
   } 
}

package main

import (
        "fmt"
        "os"
        "strconv"
)

func main() {
      args := os.Args[1:]
      if len(args) != 1 {
         fmt.Println("Please enter only one number")
         return
      }

      num, err := strconv.Atoi(args[0])
      if err != nil {
         fmt.Println("You are allowed to enter only integer number, otherwise go to hell", args[0])
         return
      }  


//     num := 123
     a := make(chan int)
     b := make(chan int)

     go calSqare(num,a)
     go calCube(num, b)

    var sum1  int
    sum1 =  <-a
    sum2 :=  <-b
    fmt.Println("Total is", sum1 + sum2)
    fmt.Printf("Type of sum1 is %T\n", sum1)
}

func calSqare(num int, a chan int) {
     sum := 0
     for num != 0 {
         digit := num %10 
         sum += digit * digit
         num /= 10
     }

     a <- sum
}

func calCube(num int, b chan int) {
     sum := 0
     for num != 0 {
         digit := num %10 
         sum += digit * digit * digit
         num /= 10
     }
 
     b <- sum
}
          


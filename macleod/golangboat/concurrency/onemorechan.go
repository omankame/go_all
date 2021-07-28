package main

import (
        "fmt"
)

func main() {
     num := 555
     sqchan := make(chan int)
     cuchan := make(chan int)

     go sqareCal(num, sqchan)
     go cuchanCal(num, cuchan)

     num1 := <-sqchan
     num2 := <-cuchan
     fmt.Println("Total is", num1 + num2)

}

func sqareCal(num int, sqchan chan int) {
     sum := 0
     dchan := make(chan int)
     go digitCal(num, dchan)        
     for v := range dchan {
         sum += v * v
     }
     sqchan <- sum

}

func cuchanCal(num int, cuchan chan int) {
     sum := 0
     dchan := make(chan int)
     go digitCal(num, dchan)
     for v := range dchan {
         sum += v * v * v
     }
     cuchan <- sum
}

func digitCal(num int, dchan chan int) {
     for num != 0 {
         digit := num % 10
         dchan <- digit
         num /= 10
     }

    close(dchan)
}

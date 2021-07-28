package main

import (
        "fmt"
)

func main() {
     ch := make(chan int)
     go dataProduce(ch)
     
     for {
         v , ok := <-ch
         if ok != true {
               break
         }
        
       fmt.Println("the value is ", v)
     }
}

func dataProduce(ch1 chan int) {
     for i := 0; i < 10; i++ {
         ch1 <- i
      }
     close(ch1)
}

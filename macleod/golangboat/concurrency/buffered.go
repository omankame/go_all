package main

import (
        "fmt"
        "time"
)

func main() {
     ch := make(chan int, 2)
     go write(ch)
     time.Sleep(2 * time.Second)
     for v := range ch {
         fmt.Println("read value", v, "from ch")
         time.Sleep(2 * time. Second)
}
}
func write(ch chan int) {
     for i :=0; i < 5; i++ {
      ch <- i
      fmt.Println("sucessfully wrote", i, "to ch")
      }
    close(ch)
}
      

package main

import (
        "fmt"
        "time"
)

func main() {
     output1 := make(chan string)
     output2 := make(chan string)

     go server1(output1)
     go server2(output2)

     
     time.Sleep(1 * time.Second)
         select {
            case v1 := <-output1:
                 fmt.Println(v1)
           
            case v2 := <-output2:
                 fmt.Println(v2)
          }
      
}

func server1(op1 chan string) {
     op1 <- "server1 data processed"
}
   
func server2(op2 chan string) {
     op2 <- "server2 data processed"
}


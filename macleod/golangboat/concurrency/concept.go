package main

import (
        "fmt"
        "time"
)

func main() {
      fmt.Println("GOroutine Tutorial")
     
      compute(5)
      compute(5)

//     var input string
//     fmt.Scanln(&input)
}

func compute(value int) {
     for i:=0; i < value; i++ {
       fmt.Println(i)
       time.Sleep(1 * time.Second)
    }
}

package main

import (
        "fmt"
        "time"
)

func main() {
     done := make(chan bool)
     fmt.Println("main going to call hello")
     go hello(done)
     <-done        //blocked untill get data received
     fmt.Println("Main received data")
}

func hello(done chan bool) {
     fmt.Println("hello going to sleep")
     time.Sleep(4 * time.Second)
     fmt.Println("hello go routine awake")
     done <- true
}


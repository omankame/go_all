package main

import (
        "fmt"
        "math/rand"
)

func main() {
     values := make(chan int)
     go calculate(values)

     value :=  <-values
     fmt.Println("Generated Value is", value)
}

func calculate(values chan int) {
     value := rand.Intn(5)
     values <- value
}


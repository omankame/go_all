package main

import (
        "fmt"
        "os"
)

func main() {
     fl, err := os.Create("Test.txt")
     if err != nil {
     fmt.Println(err)
     return
     }

    defer fl.Close()

    n, err :=  fl.WriteString("Hello World! \n") 
    if err != nil {
       fmt.Println(err)
    }
 
    fmt.Println("The n string are sucessfully written", n)

}       
    


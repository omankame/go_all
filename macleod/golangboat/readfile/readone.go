package main

import (
        "fmt"
        "io/ioutil"
)

func main() {
     data, err := ioutil.ReadFile("test.txt")
     if err != nil {
        fmt.Println(err)
        return
     }

     fmt.Println("The contents of the file:", string(data))

}

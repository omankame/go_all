package main

import (
         "fmt"
         "io/ioutil"
)

func main() {
     data, err := ioutil.ReadFile("/home/macleod/golangboat/readfile/test.txt")
     if err != nil {
     fmt.Println(err)
     return
     }

     fmt.Println("The content of the file:", string(data))
}

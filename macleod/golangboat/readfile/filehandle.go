package main

import (
         "fmt"
         "github.com/gobuffalo/packr/v2"
)

func main() {
     box := packr.New("filebox", "/home/macleod/golangboat/readfile") 
     data, err := box.FindString("test.txt")
     if err != nil {
     fmt.Println(err)
     return
     }

    fmt.Print(data)
}

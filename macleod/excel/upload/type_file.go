package main

import (
       "fmt"
       "os"
       "net/http"
)

func main() {
//     f , _ := os.Open("hello.pdf")
       f, _ := os.Open("Book99.xlsx")
     cType, err := getTypeFile(f)
     if err != nil {
     panic(err)
     }
     fmt.Println("The Content Type of given file is:", cType)
     
}
func getTypeFile(out *os.File) (string, error) {

     buffer := make([]byte, 512)
     
     _, err := out.Read(buffer)
     if err != nil {
     return "", err
     }

    contentType :=  http.DetectContentType(buffer)
    return contentType, nil
}
      
//important url https://golangcode.com/get-the-content-type-of-file/

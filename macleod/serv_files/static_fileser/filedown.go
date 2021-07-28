package main

import (
        "os"
//        "log"
        "net/http"
        "fmt"
        "io"
)

func main() {

      args := os.Args[1:]
      if len(args) != 2 {
      fmt.Println("usage: download url filename")
      os.Exit(1)
      }

      url := args[0]
      name := args[1]
      err := Downloadfile(url, name) 
      if err != nil {
      panic(err)
      }
}

func Downloadfile(url string, filepath string) error {

     out, err := os.Create(filepath)
     if err != nil {
        return err
     }
     
     defer out.Close()

     resp, err := http.Get(url)
     if err != nil {
     return err
     }
 
     defer resp.Body.Close()

     _, err = io.Copy(out, resp.Body) 
     if err != nil {
     return err
     }

     return nil      

}



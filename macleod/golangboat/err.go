package main

import (      
        "fmt"
        "os"
)

func main() {
          f, err := os.Open("map123.go")
          if err != nil {
             fmt.Println(err)
             return
          }
          fmt.Println(f.Name(), "Open Sucessfully.")
          defer f.Close() 
}

package main

import (
         "fmt"
         "os"
)

func main() {
     f, err := os.Open("map123.go")
     if err != nil {
            pErr := err.(*os.PathError)
            fmt.Println("Failed to open file at path", pErr.Path)
            return
     }

     fmt.Println(f.Name(), "Open Sucessfully")
     f.Close()
}

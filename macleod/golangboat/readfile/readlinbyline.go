package main

import (
         "fmt"
         "bufio"
         "flag"
         "os"
)


func main() {
     fptr := flag.String("fpath", "test.txt", "usage of fpath")
     flag.Parse()

     file, err := os.Open(*fptr)
     if err != nil {
        fmt.Println(err)
        return
     }

     defer func() {
        if err := file.Close(); err != nil {
           fmt.Print(err)
       }
     }()

   scanr := bufio.NewScanner(file)
   for scanr.Scan() {
       fmt.Println(scanr.Text() )
   }
}     

       

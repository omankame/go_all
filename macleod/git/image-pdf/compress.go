package main

import (
       "fmt"
       "os"
       "image/jpeg"
       "github.com/nfnt/resize"
)

func main() {
    args := os.Args[1:]
    if len(args) != 1 {
       fmt.Println("Please enter the jpg file name which you want to compress.")
       return
    }
    i_name := args[0]
    f, err := os.Open(i_name) 
    if err != nil {
    fmt.Println(err)
    return
    }

    defer f.Close()

    img, _ := jpeg.Decode(f)
    
    m := resize.Resize(1000, 0, img , resize.Lanczos3)

    out, _ := os.Create("compres" + i_name)

    jpeg.Encode(out, m, nil) 

}       


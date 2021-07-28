package main

import (
         "fmt"
//         "io"
//         "io/ioutil"
         "os"
         "path/filepath"    
         
)

var srcfd *os.File

var dstfd *os.File



func main() {
     args := os.Args[1:]
     if len(args) != 2 {
        fmt.Println("Please enter source file name/path and dst path")
        return
     }
     var src, dst string

     src =   args[0]
     dst,_ = filepath.Abs(args[1]) // absolute path representation else current file
     
     srcf := filepath.Base(src)     

     dstf := filepath.Join(dst, srcf)
//     srcb, _  :=  ioutil.ReadFile(src)  //reading all data as []byte slice
   
//     ioutil.WriteFile( dstf, srcb, 0664) 
   
     fmt.Println(src, dst, dstf, srcf)
           

}  

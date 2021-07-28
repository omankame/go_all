package main

import (
         "fmt"
//         "io"
         "io/ioutil"
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

     

     src = args[0]
     srcf := filepath.Base(src) // for geeting last argument of source file path 
     fmt.Println(src)
     fileInfo, err := os.Stat(src)
     if err != nil {
        fmt.Println(err)
     }

     if fileInfo.IsDir() {
        fmt.Println("File can not open")
        return
     }

    
 
     dst,_ = filepath.Abs(args[1]) // absolute path representation else current file 
     dstf := filepath.Join(dst, srcf) 

     srcb, err  :=  ioutil.ReadFile(src)  //reading all data as []byte slice
     if err != nil {
        fmt.Println(err)
        return
     }
     
     ioutil.WriteFile( dstf, srcb, 0664) 
   
     fmt.Println(src, dst, srcf, dstf)
           

}  

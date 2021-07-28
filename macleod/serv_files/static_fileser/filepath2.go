// https://progolang.com/listing-files-and-directories-in-go/

package main

import (
        "fmt"
        "path/filepath"
        "os"
)

func main() {

    WalkAllFilesInDir("/home/macleod/serv_files")
}

func WalkAllFilesInDir(dir string) error {
     return filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
     if e != nil{
     return e
     }

     if info.Mode().IsRegular() {
        fmt.Println("file name is :", info.Name())
     //   fmt.Println("file path is :", path)
     }
     return nil
     })

}  
     

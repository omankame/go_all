// https://progolang.com/listing-files-and-directories-in-go/  

package main

import (
        "fmt"
        "path/filepath"
)

func main() {

     list, err := filepath.Glob("/home/pdf/*.pdf")
     if err != nil {
     fmt.Println(err)
     return
     }
     for _, v := range list {
     fmt.Println(v)
     }
 

}

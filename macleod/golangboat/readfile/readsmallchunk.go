package main

import (
        "fmt"
        "bufio"
        "flag"
        "os"
        
)

func main() {
     fdata := flag.String("fpath", "test.txt", "use - and specify the file")    
     flag.Parse()
 
     file, err := os.Open(*fdata)
     if err != nil {
     fmt.Println(err)
     return
     }

     defer file.Close()

     r := bufio.NewReader(file)
     dchunk := make([]byte, 3)

     for {
       n, err := r.Read(dchunk) 
       if err != nil {
          fmt.Println("Error reading file.", err)
          break
        }
       fmt.Println(string(dchunk[0:n]))
    }

}

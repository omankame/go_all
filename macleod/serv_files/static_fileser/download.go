package main

import (
        "os"
        "io"
        "net/http"
        "fmt"
        "github.com/dustin/go-humanize"
        "strings"

)
type WordCounter struct {
     Total  uint64
}

func (wc *WordCounter) Write(p []byte) (int,  error) {
      n := len(p)
      wc.Total += uint64(n)
      wc.printProgress()
      return n, nil  
} 

func (wc WordCounter) printProgress() {
     fmt.Printf("\r%s", strings.Repeat(" ", 50))

     fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func main() {
     args := os.Args[1:]
     if len(args) != 2 {
     fmt.Println("usage: download url filename")
     os.Exit(1)
     }

     url := args[0]
     filename := args[1]
    
     err := downloadFile(url, filename )
     if err != nil {
     fmt.Println(err)
     }

}

func downloadFile(url string, filepath string) error {

     file, _ := os.Create(filepath + ".tmp") 

     defer file.Close()

     resp, err := http.Get(url) 
     if err != nil {
     return err
     }

     defer resp.Body.Close()

     counter := &WordCounter{}
     _, err = io.Copy(file, io.TeeReader(resp.Body, counter))
     if err != nil {
        return err
       }

    fmt.Println()

    err = os.Rename(filepath+".tmp", filepath)
    if err != nil {
        return err
    } 
    
    return nil

} 
    


package main

import (
        "fmt"
        "net/http"
        "os"

)


var sites = []string{
	"https://github.com",
	"https://google.com",
	"https://stackoverflow.com",
	"https://facebook.com",
	"https://twitter.com",
	"https://golang.org",
	"https://forum.golangbridge.org",
	"https://packtpub.com/",
}

func main() {
     args := os.Args[1:]
     fmt.Println(args[0])
     switch args[0]  {
 
     case "seq":
           get()
     case  "conc": 
            getConcurrent()
     default:
        fmt.Println("PLease type 'seq' or 'conc'.")

     }

}

func get() {

     for _, v := range sites {
        resp, err :=  http.Get(v)
        if err != nil {
           fmt.Println(err)
           return
        }
        fmt.Printf("%s %d \n", v, resp.StatusCode)
        resp.Body.Close()
     }
}

func getConcurrent() {
     ch := make(chan string)

     for _, v := range sites {
      
         go func(v string) {
            resp, err := http.Get(v)
            if err != nil {
            ch <- fmt.Sprintf("%s", err)
            return
           }

          ch <- fmt.Sprintf("%s %d \n", v, resp.StatusCode)
          resp.Body.Close()
          }(v)
      }

      for range sites {
          fmt.Println(<-ch)
      }

}      
     
           

       

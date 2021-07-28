package main 

import (
//         "fmt"
         "os"
         "log"
         "text/template"
       )

func main() {
         
        tpl, err := template.ParseFiles("tpl.gohtml")
        if err != nil {
           log.Fatal(err)
        }

       nf, err := os.Create("index.html")
       if err != nil {
           log.Fatal(err)
        }
        defer nf.Close()
        tpl.Execute(nf, nil) 
}


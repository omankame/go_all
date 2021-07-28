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
          os.Exit(1)
         }
         
         tpl.Execute(os.Stdout, nil)




}

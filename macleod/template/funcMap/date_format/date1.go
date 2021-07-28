package main

import (
         "os"
         "text/template"
         "log"
         "time"
       )

var tpl *template.Template

 
func main() {
 
      t := time.Now()
      tpl = template.Must(template.ParseFiles("tpl1.gohtml"))       
 
      err := tpl.Execute(os.Stdout,  t)
      if err != nil {
         log.Fatal(err)
         os.Exit(1)
     }

}

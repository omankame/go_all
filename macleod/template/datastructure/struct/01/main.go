package main

import (
         "os"
         "log"
         "text/template"
       )

var tpl *template.Template

type sage struct {
     Name    string
     Motto   string
}

func init() {
      tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

     buddha := sage {
           Name:   "Buddha",
           Motto:  "The belief of no beliefs",
     }

    err :=  tpl.Execute(os.Stdout, buddha)
    if err != nil {
       log.Fatal(err)
        os.Exit(1)
    }

}
     

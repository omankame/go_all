package main

import ( 
         "os"
         "text/template"
         "log"
)

var tpl *template.Template

func main() {

     str := []string{"zero", "one", "two", "three", "four", }
    tpl = template.Must(template.ParseFiles("tpl.gohtml"))
    
    err := tpl.Execute(os.Stdout, str)
    if err != nil {
       log.Fatal(err)
     os.Exit(1)
    }

}

package main

import  (
          "os"
          "log"
          "text/template"
        )
var tpl *template.Template

func init () {
     tpl = template.Must(template.ParseFiles("tpl.gohtml"))
} 
 

func main() {

       err :=   tpl.Execute(os.Stdout, 42)
       if err != nil {
          log.Fatal(err)
       }
       
      err =   tpl.Execute(os.Stdout, "SatNam")
       if err != nil {
          log.Fatal(err)
       }


}


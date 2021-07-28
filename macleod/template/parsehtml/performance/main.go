package main

import (
        "log"
        "text/template"
        "os"
       )
var tpl *template.Template

func init() {
       tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
       err := tpl.Execute(os.Stdout, nil)
       if err != nil {
          log.Fatal(err)
       }      

       tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml" , nil)
       tpl.ExecuteTemplate(os.Stdout, "two.gohtml" , nil)
       tpl.ExecuteTemplate(os.Stdout, "one.gohtml" , nil)

}

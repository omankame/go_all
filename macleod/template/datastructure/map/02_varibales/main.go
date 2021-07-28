package main

import (
         "os"
         "log"
         "text/template"
       )

var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

        sage := map[string]string {
                "India":    "Gandhi",
                "America":  "MLK",
                "Meditate": "Buddha",
                "Love":     "Jesus",
                "Prophet":  "Muhammad" }
                     

       err := tpl.Execute(os.Stdout, sage)
       if err != nil {
       log.Fatal(err)
       os.Exit(1)
       }

}

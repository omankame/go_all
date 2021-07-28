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

         sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Mohmmad"}
         err := tpl.Execute(os.Stdout, sages)
         if err != nil {
            log.Fatal(err)
            os.Exit(1)
          }

}

    

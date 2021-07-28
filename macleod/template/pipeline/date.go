package main

import (
         "os"
         "log"
         "text/template"
         "time"
)

var tpl *template.Template

var fm = template.FuncMap {
    "fdateMDY": formatdate,
}

func formatdate(t time.Time) string {
     return t.Format("02-01-2006")
}

func init() {
     tpl = template.Must(template.New("timee").Funcs(fm).ParseFiles("date.gohtml"))
}

func main() {

    t := time.Now()
     err := tpl.ExecuteTemplate(os.Stdout, "date.gohtml", t)
     if err != nil {
        log.Fatal(err)
       os.Exit(1)
     }

}
 

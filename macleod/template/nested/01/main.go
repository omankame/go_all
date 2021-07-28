package main

import (
        "os"
        "text/template"
        "log"
)

var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseGlob("template/*.gohtml"))
}

func main() {

     err := tpl.ExecuteTemplate(os.Stdout,"index.gohtml", 42)
     if err != nil {
        log.Fatal(err)
    }

}


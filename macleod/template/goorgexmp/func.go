package main

import (
         "log"
         "text/template"
         "strings"
         "os"
       )

const templateText = `
Input: {{printf "%q" .}}
Output 0: {{title .}}
Output 1: {{title . | printf "%q"}}
Output 2: {{printf "%q" . | title}}
`

func main() {

      var fm = template.FuncMap {
          "title": strings.Title,
      }

//      var tpl *template.Template  effective if you are accesing file from outside
        
        tmpl, err := template.New("").Funcs(fm).Parse(templateText)
        if err != nil {
           log.Fatal(err)
       }
       
       err = tmpl.Execute(os.Stdout, "the go programming language")  

        


}


package main

import (
          "os"
          "log"
          "text/template"
       )

 
func main() {
 
           tpl, err := template.ParseGlob("templates/*")
           if err != nil {
           log.Fatal(err)
           }

           err = tpl.Execute(os.Stdout, nil)
           if err != nil {
           log.Fatal(err)
           }

           tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml" , nil) 
           tpl.ExecuteTemplate(os.Stdout, "two.gohtml" , nil) 
           tpl.ExecuteTemplate(os.Stdout, "one.gohtml" , nil) 
           tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml" , nil) 
}

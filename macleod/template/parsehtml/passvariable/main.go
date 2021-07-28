package main

import (
         "log"
         "os"
         "text/template"
//         "io"
       )

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseFiles("tpl.gohtml"))

}

func main() {
     err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Release self-focus; embrace other-focus.`)
     if err != nil {
     log.Fatal(err)
     os.Exit(1)
     }
     
     nf, err :=  os.Create("index.gohtml")
     if err != nil {
     log.Fatal(err)
     os.Exit(1)
     }

     defer nf.Close() 
 
//     io.Copy(nf, strings.NewReader(str))
     err = tpl.Execute(nf,  `Release self-focus; embrace other-focus.`)
     if err != nil {
     log.Fatal(err)
     os.Exit(1)
     }


}      
   
  

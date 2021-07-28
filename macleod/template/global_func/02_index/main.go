package main

import ( 
         "os"
         "text/template"
         "log"
       )

var tpl *template.Template 

func init() {
     tpl = template.Must(template.ParseFiles("tpl.gohtml"))

}

func main() {

       xl := []string{"zero", "one", "two", "three", "four", "five"}

       data := struct {
         Words  []string
         Lname	 string
       }{
           xl,   
           "Macleod",
        }

        err := tpl.Execute(os.Stdout, data)
        if err != nil {
           log.Fatal(err)
         os.Exit(1)
        }

}

package main

import (
        "cms"
        "os"
)

func main() {

     p := &cms.Page {
           Title: "Hello World!",
           Content: "This is the body of our page",
      }

      cms.Tpl.ExecuteTemplate(os.Stdout, "index", p)

} 


         

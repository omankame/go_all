package main

import (
//       "fmt"
       "text/template"
       "os"
)

const templ = `
Hi {{.Name}}!
Welcome {{.Place}}.
Please bring {{.ToBring}}
`

func main() {

//     tpl := template.New("templ")
  
     var data = map[string]string {
         "Name": "BOB",
         "Place": "Home",
         "ToBring":  "some beers",
      }
    
     tpl, _ := template.New("").Parse(templ)
     tpl.Execute(os.Stdout, data) 

    data["Name"] = "Alice"
    data["ToBring"] = "a Teddy Bear"

    tpl.Execute(os.Stdout, data)
     
}
         




// o/p ==>

// Hi Bob!
// Welcome Home.
// Please bring some beers

// Hi Alice!
// Welcome Home.
// Please bring a Teddy Bear

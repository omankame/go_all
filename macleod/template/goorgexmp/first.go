//

package main

import (
	"log"
	"os"
	"text/template"
//        "fmt"
)

func main() {
	// Define a template.
	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`
  
//          fmt.Println(letter)

type Receipent struct {
     Name	string
     Gift	string
     Attended   bool
}
     var receipents = []Receipent{
           {"Aunt Mildred", "bone china tea set", true},
           {"Uncle John", "moleskin pants", false},
	   {"Cousin Rodney", "", false},
          }
       
     var tpl *template.Template

     tpl = template.Must(template.New("letter").Parse(letter))

     for _, r := range receipents { 
         err := tpl.Execute(os.Stdout, r)
        if err != nil {
         log.Fatal(err)
         os.Exit(1)
        }
     }


}




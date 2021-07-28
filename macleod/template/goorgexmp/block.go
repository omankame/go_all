package main 

import (
         "os" 
         "log"
         "text/template"
         "strings"
         "fmt"
)

const (
		master  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		overlay = `{{define "list"}} {{join . ", "}}{{end}} `
	)

//var tpl *template.Template

var fm = template.FuncMap {
         "join": strings.Join,
}

func main() {

       guardians := []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
 
       masterTpl, err := template.New("master").Funcs(fm).Parse(master)
       if err != nil {
        log.Fatal(err)
        os.Exit(1)
       }

       cloneTpl, err := template.Must(masterTpl.Clone()).Parse(overlay)
       if err != nil {
        log.Fatal(err)
        os.Exit(1)
       }
  
       err = masterTpl.Execute(os.Stdout, guardians)
       if err != nil {
        log.Fatal(err)
        os.Exit(1)
       }

       err = cloneTpl.Execute(os.Stdout, guardians)
        if err != nil {
        log.Fatal(err)
        os.Exit(1)
       }

     fmt.Println()
}

          
         



package cms

import (
         "html/template"
        
)

var Tpl *template.Template

func init() {
     Tpl = template.Must(template.ParseGlob("../templates/*"))
                                            
}

type Page struct {
     Title 	string
     Content 	string
}



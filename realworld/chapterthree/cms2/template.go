package cms2

import (
        "html/template"
        "time"
)

var Tpl *template.Template

func init() {
     Tpl = template.Must(template.ParseGlob("templates/*"))
}

// Page is the struct used for each webpage
type Page struct {
     Title 	string
     Content	string
     Posts	[]*Post
}

// Post is the struct used for each blog post
type Post struct {
     Title 	string
     Content	string
     DatePublished  time.Time
     Comments	[]*Comment
}

// Comment is the struct used for each comment

type Comment struct {
     Author	string
     Comment	string
     DatePublished  time.Time
}


      

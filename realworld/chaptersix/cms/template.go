package cms

import (
        "html/template"
        "time"

)

var Tpl *template.Template

func init() {
     Tpl = template.Must(template.ParseGlob("templates/*"))
}

type Page struct {
        ID      int
        Title   string
        Content string
        Posts   []*Post
}

type Post struct {
        ID            int
        Title         string
        Content       string
        DatePublished time.Time
        Comments   []*Comment
}

type Comment struct {
        ID            int
        PostID        int
        Author        string
        Comment       string
        DatePublished time.Time
}

 

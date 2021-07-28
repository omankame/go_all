package cms2

import (
         "net/http"
         "time"
         "strings"
         "fmt"
)

func ServeIndex(w http.ResponseWriter, req *http.Request) {
     p := &Page {
             Title:  "GO PROJECT CMS",
             Content: "Welcome to our Home Page!",
             Posts: []*Post { 
                     &Post{
                       Title: "Hello, World!",
                       Content: "Hello World! Thanks for coming to the site.",
                       DatePublished: time.Now(),
                     },
             
                     &Post{
                       Title: "A Post with Comments",
                       Content: "Here is controversial commanet. It will attract the comments",
                       DatePublished: time.Now().Add(-time.Hour),
                       Comments:  []*Comment {
                                     &Comment{
                                         Author: "Onkar Mankame",
                                         Comment: "I just commented on my Posts... ",
                                         DatePublished: time.Now().Add(-time.Hour / 2),
                                     },
                         },
                    },
          },       
    }
     
   Tpl.ExecuteTemplate(w, "page", p)         
                 
}

func ServePage(w http.ResponseWriter, req *http.Request) {
     path := strings.TrimLeft(req.URL.Path, "/page/")
     fmt.Printf("%[1]T %[1]v\n", path)
     
     if path == "" {
           http.NotFound(w,req)
           return
      }     

     p := &Page {
           Title: strings.ToTitle(path),
           Content: "Here is my page",
         }

    Tpl.ExecuteTemplate(w, "page", p)
}
     

func ServePost(w http.ResponseWriter, req *http.Request) {
      path := strings.TrimPrefix(req.URL.Path, "/post/")
      fmt.Printf("%[1]T %[1]v\n", path)
     if path == "" {
        http.NotFound(w, req)
        return
     } 
    
     p := &Post {
           Title: strings.ToTitle(path),
           Content: "Here is my page with Post",
           Comments: []*Comment {
                         &Comment {
                          Author: "Onkar Post",
                          Comment: "Its complex, It need Practice",
                          DatePublished: time.Now().Add(-time.Hour),
                         },
                    },
            }

       Tpl.ExecuteTemplate(w, "post", p)

}  

func HandleNew(w http.ResponseWriter, req *http.Request) {
     switch req.Method {
   
        case "GET":
              Tpl.ExecuteTemplate(w, "new", nil)

        case "POST":
             title := req.FormValue("title")
             content := req.FormValue("content")
             contentType := req.FormValue("content-type")
             req.ParseForm()

             if contentType == "page" {
                    Tpl.ExecuteTemplate(w, "page",   &Page {
                              				Title: title,
                          				Content: content,
                                                       })
                     return
              }

              if contentType == "post" {
                     Tpl.ExecuteTemplate(w, "post", &Post {
                                                        Title: title,
                                                        Content: content,
                                                       })

                     return
               }  
  

        default:
        http.Error(w, "Method not supported: "+req.Method, http.StatusMethodNotAllowed)
    }
}

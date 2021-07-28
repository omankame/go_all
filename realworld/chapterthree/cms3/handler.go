package cms3

import (
          "fmt"
          "net/http"
          "strings"
          "time"
)

func ServeIndex(w http.ResponseWriter, req *http.Request) {
               p := &Page{
                          Title: "Go Projects CMS",
                          Content: "Welcome to our home page!",
                          Posts: []*Post {      
                                   &Post {
                                          Title:   "Hello, World!",
                                          Content: "Hello world! Thanks for coming to the site.",
                                          DatePublished: time.Now(),
                                    },
                  
                                   &Post {
                                          Title:   "A Post with Comments",
                                          Content: "Here's a controversial post. It's sure to attract comments.",
                                          DatePublished: time.Now().Add(-time.Hour),
                                          Comments: []*Comment {
                                                       &Comment {
                                                                 Author: "Onkar Mankame",
								 Comment: "Haaha, I just commented on my Post",
 								 DatePublished: time.Now().Add(-time.Hour),
                                                       },
                                          },
                                    },
                            },
                   }

            Tpl.ExecuteTemplate(w, "page", p)

}

func ServePage(w http.ResponseWriter, req *http.Request) {
     path := strings.TrimLeft(req.URL.Path, "/page/")
     
     if path == "" {
             pages, err := GetPages()
             if err != nil {
                     http.Error(w, err.Error(), http.StatusInternalServerError)
                     return
             }
         Tpl.ExecuteTemplate(w, "pages.gohtml", pages)
         return                     

     }


     page, err := GetPage(path) 
           if err != nil {
              http.Error(w, err.Error(), http.StatusNotFound)
              return
           }

    Tpl.ExecuteTemplate(w, "page", page)	

}


func ServePost(w http.ResponseWriter, req *http.Request) {
     path := strings.TrimLeft(req.URL.Path, "/post/")
    
     if path == "" {
             http.NotFound(w, req)
             return
     }

     p := &Post{
           Title: strings.ToTitle(path),
           Content: "Here is my page",
           Comments: []*Comment {
                       &Comment {
                           Author:  "Onkar Mankame",
                           Comment: "Looks Great!",
                           DatePublished: time.Now(),
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
                     p := &Page {
                                  Title:   title,
				Content: content,
			}
                     _, err := CreatePage(p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			Tpl.ExecuteTemplate(w, "page", p)
			return
		  }
 
                  if contentType == "post" {
                             p := &Post {
                                  Title:         title,
				Content:       content,
				DatePublished: time.Now(),
			       }
			     id, err := CreatePost(p)
			       if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
                              Tpl.ExecuteTemplate(w, "post", p)
			     fmt.Printf("Created post with id: %d\n", id)
			      return

                  }

               default: 
                 http.Error(w, "Method not supported: "+req.Method, http.StatusMethodNotAllowed)              
                    
            }

}
           

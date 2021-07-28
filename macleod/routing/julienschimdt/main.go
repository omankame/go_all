package main

import (
//        "os"
        "fmt"
        "html/template"
        "net/http"
//        "log"
//        "io"
          "github.com/julienschmidt/httprouter"
)
var tpl *template.Template

func init() {

     tpl = template.Must(template.ParseGlob("template/*"))  
}

func main() {

//    nf, _ := os.Create("new.route")    
 
//    tpl.ExecuteTemplate(wr io.Writer, name string, data interface{}) 
//     tpl.ExecuteTemplate(os.Stdout, "about.gohtml" , nil)

      mux := httprouter.New()
      mux.GET("/", index)      
      mux.GET("/about", about)
      mux.GET("/contact", contact)
      mux.GET("/apply", apply)
      mux.POST("/apply", postapply)
      mux.GET("/user/:name", Hello)
      mux.GET("/blog/:category/:article", blogRead)
//      mux.POST("/blog/:category/:article", blogWrite)


      http.ListenAndServe(":8080", mux)

}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
    }

func blogRead(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
         fmt.Fprintf(w, "READ CATEGORY, %s!\n", ps.ByName("category"))
        fmt.Fprintf(w, "READ ARTICLE, %s!\n", ps.ByName("article"))
}


func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
           tpl.ExecuteTemplate(w, "index.gohtml" , nil)
}

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
           tpl.ExecuteTemplate(w, "about.gohtml" , nil)
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
           tpl.ExecuteTemplate(w, "contact.gohtml" , nil)
}

func apply(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
           tpl.ExecuteTemplate(w, "apply.gohtml" , nil)
}

func postapply(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
           tpl.ExecuteTemplate(w, "applyProcess.gohtml" , nil)
}

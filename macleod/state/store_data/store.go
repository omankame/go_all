package main

import (
        "os"
        "path/filepath"
        "io/ioutil"
        "html/template"
        "net/http"
        "fmt"
)

var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseGlob("templates/*"))
}


func main() {
     http.HandleFunc("/", foo)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {

     var s string
     fmt.Println(req.Method)
     if req.Method == http.MethodPost {
   
     f, h, err := req.FormFile("q")
     if err != nil {
     http.Error(w, err.Error(), http.StatusInternalServerError)
     return
     }
   
     fmt.Println("File :", f , "\nHeader :", h, "\nError :", err)
     
     bs, _ := ioutil.ReadAll(f)

     s = string(bs)     

     dst, _ := os.Create(filepath.Join("./user", h.Filename))
     defer dst.Close()  
     _, _ = dst.Write(bs)
     
     
     }
     w.Header().Set("Content-Type", "text/html; charset=utf-8") 
     tpl.ExecuteTemplate(w , "index.gohtml", s) 

}

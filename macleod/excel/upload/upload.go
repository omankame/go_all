package main

import (
         "fmt"
         "net/http"
         "html/template"
         "io/ioutil"
         "os"
)

var tpl *template.Template

func init() {
     tpl  = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
      
     http.HandleFunc("/", wel)
     http.HandleFunc("/upload", Upload)
     http.ListenAndServe(":8080", nil) 
}

func wel(w http.ResponseWriter, req *http.Request) {
     fmt.Fprint(w, "Welcome to UPLOAD Demo")
}

func Upload(w http.ResponseWriter, req *http.Request) {
     var s string 
     if req.Method == http.MethodPost {
        f, h, err := req.FormFile("myFile")
        if err != nil {
        http.Error(w , "SOmething Weird happened", http.StatusInternalServerError)
        }
        fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)  
        defer f.Close()
        
        nf, _ := os.Create(h.Filename) 
       
        bs, _ := ioutil.ReadAll(f)
        nf.Write(bs)
        defer nf.Close()
//        s = string(bs)
//        fmt.Println(s)
        }
      
//     application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
     w.Header().Set("Content-Type", "text/html; charset=utf-8")
     tpl.ExecuteTemplate(w, "index.gohtml", s)
}

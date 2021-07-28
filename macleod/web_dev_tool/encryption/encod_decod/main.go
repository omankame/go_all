package main

import (
         "fmt"
         "net/http"
         "html/template"
         "encoding/base64"
)

var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseGlob("template/*.gohtml"))
}

func main() {
     fmt.Println("Starting Web Server")
     http.HandleFunc("/", index)
     http.HandleFunc("/encode", encode)
     http.HandleFunc("/decode", decode)
     http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {


     tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
func encode(w http.ResponseWriter, req *http.Request) {
     if req.Method == "POST" {
                req.ParseForm()
                text := req.FormValue("w3review")
//                fmt.Println(text)
                str := base64.URLEncoding.EncodeToString([]byte(text))
//                fmt.Println(str)
                tpl.ExecuteTemplate(w, "encode.gohtml", str)
                return
     }

     tpl.ExecuteTemplate(w, "encode.gohtml", nil)
}
func decode(w http.ResponseWriter, req *http.Request) {
     if req.Method == "POST" {
                req.ParseForm()
                text := req.FormValue("w3review")
                stbyte, err := base64.URLEncoding.DecodeString(text) 
                if err != nil {
                   fmt.Println(err)
                   return
                }
//                fmt.Println(string(stbyte))
                str := string(stbyte)
                tpl.ExecuteTemplate(w, "decode.gohtml", str)
                return
     }


     tpl.ExecuteTemplate(w, "decode.gohtml", nil)
}

// vimp  if space exists in between <textarea></textarea> then it will gives illegal base64 data at input byte 0 error. 

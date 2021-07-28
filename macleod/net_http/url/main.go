package main

import ( 
        "log"
        "net/http"
        "html/template"
//        "os"
        "net/url"
)

type hotdog int
var tpl *template.Template
func init() {
     tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
      err := r.ParseForm()
      if err != nil {
      log.Println(err)
      }

      data := struct {
              Method	string
              URL 	*url.URL
              Submissions url.Values
            }{
              r.Method,
              r.URL,
              r.Form,
           }
      tpl.ExecuteTemplate(w , "index.gohtml", data) 	 

}

func main() {

     var d hotdog
     http.ListenAndServe(":8080", d)

}

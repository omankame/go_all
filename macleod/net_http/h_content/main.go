package main

import ( 
         "log"
         "net/http"
         "net/url"
         "html/template"
)

var tpl *template.Template

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r  *http.Request) {

      err := r.ParseForm()
      if err != nil {
      log.Println(err)
      }
      
      data := struct {
               Method	string
               URL      *url.URL
               Header   http.Header
               Submissions map[string][]string
            }{  
              r.Method,
              r.URL,
              r.Header,
              r.Form,
          }

         tpl.ExecuteTemplate(w , "index.gohtml", data)
 
}

func init() {
     tpl = template.Must(template.ParseFiles("index.gohtml"))

}

func main() {
      var d hotdog
      http.ListenAndServe(":8080", d)
}


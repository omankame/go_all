package api

import (
        "encoding/json"
        "net/http"
//        "fmt"
        "cms"
)

func Doc(w http.ResponseWriter, req *http.Request) {

         data := map[string]string{
                     "all_pages_url":   "/pages",
	             "create_page_url": "/newpage",
	             "page_url":        "/pages/{id}",
                    }

          writeJson(w, data) 


}

func AllPages(w http.ResponseWriter, req *http.Request) {

     data, err := cms.GetPages()
     if err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
         return
     }
     writeJson(w, data)
}

func writeJson( w http.ResponseWriter, data interface{}) {
     w.Header().Set("Content-Type", "application/json; charset=utf-8")
     resJson, err := json.MarshalIndent(data, "", "\t")
     if err != nil {
        errJson(w, err.Error(), http.StatusInternalServerError)
        return
     }

     w.Write(resJson)
} 

func errJson(w http.ResponseWriter, err string, status int) {

     w.Header().Set("Content-Type", "application/json")
     w.WriteHeader(status)
     w.Write([]byte("{\n\terror: " + err + "\n}\n"))
}
     

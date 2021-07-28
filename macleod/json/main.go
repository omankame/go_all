/* vimpreference while creating test.json https://dev.to/christalib/append-data-to-json-in-go-5gbj
 https://www.golangprograms.com/golang-writing-struct-to-json-file.html#:~:text=The%20json%20package%20has%20a,a%20file%20in%20JSON%20format.&text=The%20Salary%20struct%20is%20defined,then%20serialize%20with%20the%20json.  */

// https://www.golangprograms.com/web-application-to-read-and-write-json-data-in-json-file.html must check
package main

import (
        "net/http"
        "html/template"
        "fmt"
        "encoding/json"
	"io/ioutil"
        "os"
      )
var tpl *template.Template

var filename = "test.json"

type Employee struct {
     Name   string `json="Name"`
     Gender string `json="Gender"`
     Age    string `json="Age"`
}

func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
    
     err := checkFile(filename)
     if err != nil {
        fmt.Println("Some critical error can not continue without fixing it", err)
        return
     }
}
func checkFile(fn string) error {
     _, err := os.Stat(fn)
     if os.IsNotExist(err) {
        _, err = os.Create(fn)
        if err != nil {
           return err
        }
     }
    return nil
}

func main() {
     http.HandleFunc("/", index)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
//     var emp Employee
     
     n := req.FormValue("name")
     g := req.FormValue("gender")
     a := req.FormValue("age")

     tpl.ExecuteTemplate(w , "index.gohtml", Employee{n, g, a})  

     emp := &Employee{n, g, a}

// Reading file "test.json"

     data := []Employee{}
     fdata, _ := ioutil.ReadFile(filename)
     json.Unmarshal(fdata, &data) 

     data = append(data, *emp)

// writing data to test.json
     file , _ := json.Marshal(data)
     err := ioutil.WriteFile(filename, file, 0644)
     if err != nil {
     fmt.Println("Error while creating and/or inserting data", err)
     return
     }          
}

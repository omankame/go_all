package main

import (
         "net/http"
         "github.com/gorilla/mux"
         "fmt"
         "encoding/json"
         "io/ioutil"
         "os"
)

type animalS struct {
     Name   string `json:"name"`
     Sound  string `json:"sound"`
} 

func main() {
      
       r := mux.NewRouter()

       r.HandleFunc("/", index)
       r.HandleFunc("/ping", ping)
       r.HandleFunc("/animal/{name}", animal)
       http.ListenAndServe(":3333", r)

}

func index(w http.ResponseWriter, req *http.Request) {
     fmt.Fprintln(w, "Welcome to Animal Page")
}

func ping(w http.ResponseWriter, req *http.Request) {
     fmt.Fprintln(w, "pong")
}

func animal(w http.ResponseWriter, req *http.Request) {

     var AnimalS []animalS

     file, err := os.Open("data.json") 
     if err != nil {
        fmt.Println(err)
        return
     }

     defer file.Close()


     byteFile, err := ioutil.ReadAll(file)
     if err != nil {
        fmt.Println(err)
        return
     }

     json.Unmarshal(byteFile, &AnimalS) 

//     fmt.Println(AnimalS)

   
     vars := mux.Vars(req)
     name  := vars["name"]
//     fmt.Fprint(w, name)

     for _, v := range AnimalS {
         if name == v.Name {
            fmt.Fprintln(w, "Name: "+v.Name + "\t" + "Sound: "+v.Sound)
            return
         }

     }  
     fmt.Fprintln(w, "No Animal Found")
}


//  https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e

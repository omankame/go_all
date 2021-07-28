package main

import (
        "net/http"
        "encoding/json"
        "fmt"
)

type Bird struct {
     Species    string `json: "species"`
     Description  string `json: "description"`
}

var birds []Bird

func getBirdHandler(w http.ResponseWriter, req *http.Request) {
     // Convert the "birds" variable to json
     
     birdsList, err := json.Marshal(birds)
     if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
     }

     w.Write(birdsList)

} 
 
func createBirdHandler(w http.ResponseWriter, req *http.Request) {
     bird := Bird{}

     err := req.ParseForm()
      if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
     }

     bird.Species = req.FormValue("species")
     bird.Description = req.FormValue("description")

     birds = append(birds,bird)

     http.Redirect(w, req, "/assets/", http.StatusSeeOther)

}
     
     

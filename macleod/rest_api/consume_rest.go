package main

import (
        "fmt"
        "net/http"
        "io/ioutil"
        "encoding/json"
)

type PokeMonSpecies struct {
     Name	string `json:"name"`
}

type PokeMon struct {
     Entry_no	int `json:"entry_number"`
     Species	PokeMonSpecies `json:"pokemon_species"`
}

type Response struct {
     Name	string `json"name"`
     Pokemon []PokeMon `json:"pokemon_entries"`
}  

func main() {
     url := "http://pokeapi.co/api/v2/pokedex/kanto/"   
     
     resp, err := http.Get(url) 
     if err != nil {
     fmt.Println(err)
     }

     responsedata, _ := ioutil. ReadAll(resp.Body) 
//     fmt.Println(string(body))

   
    var responseObject Response
    json.Unmarshal(responsedata, &responseObject)

    fmt.Println(responseObject.Name)
    fmt.Println(len(responseObject.Pokemon))

    for i := 0; i < len(responseObject.Pokemon); i++ {
        fmt.Println(responseObject.Pokemon[i].Species.Name)
   }

}

// https://tutorialedge.net/golang/consuming-restful-api-with-go/      




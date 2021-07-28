package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

// your JSON structure as a byte slice


func main() {

    file, err  := ioutil.ReadFile("test.json")
                  if err != nil {
                                  panic(err)
                  }

    c := make(map[string]json.RawMessage)

     e := json.Unmarshal(file, &c)

     if e != nil {
        panic(e)
     }

     fmt.Println(c)

}
     

package main

import (
        "encoding/xml"
//        "net/http"
        "fmt"
)

type Person struct {
     Name	string
     Age	int
}

func main() {
//     p := &Person {
//          Name: "onkar",
//          Age:  "31",
//     }

    p := &Person{"ONKAR", 31} 
    v, _ := xml.MarshalIndent(p, " ", " ")  
    fmt.Println(string(v))
}




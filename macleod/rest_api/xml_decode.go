package main

import (
        "encoding/xml"
        "fmt"
 
type Book struct {
     Name    string
     Author  string
}	

func main() {

s := `
        <Book>
         <Name>War And Peace</Name>
         <Author>Leo Tolstoy</Author>
        </Book>
    `
    b := &Book{}
    
    xml.Unmarshal([]byte(s), b)
 

}

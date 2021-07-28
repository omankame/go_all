package main

import (
        "fmt"
        "os"
        
)

func main() {
     fl, err := os.Create("onkar.txt")
     if err != nil {
        fmt.Println(err)
        return
     }

     defer fl.Close()
     sl := []string{"Hello GO World!", "Go IS One Of the best", "GO Goaaa Goooo",}
     
     for _, v := range sl {
         fmt.Fprintln(fl, v)
       
     }          
}

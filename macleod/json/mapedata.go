package main

import (
      "fmt"
)

type M map[string]interface{}
func main() {
     var data = M{"Age": "28", "Gender": "female", "Name": "Uttara" }

     fmt.Println(data["Name"])
}
     

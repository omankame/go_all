package main 

import "fmt"
 
func main() {
   
     sages := map[string]string{
               "India":    "Gandhi",
                "America":  "MLK",
                "Meditate": "Buddha",
                "Love":     "Jesus",
                "Prophet":  "Muhammad"}

     fmt.Println(sages)
     for k, v := range sages {
         fmt.Printf("%s - %s \n", k, v)
       }

}
    

package main

import (
        "github.com/stmcginnis/gofish"
        "fmt"
//        "os"
)

func main() {

     config := gofish.ClientConfig {
               Endpoint: "https://10.88.32.70",
               Username: "hpadmin",
               Password: "5g_aci_36s74s5k1sc_no_os",
               Insecure: true,
               BasicAuth: true,     
              }

     c, err := gofish.Connect(config)
     if err != nil {
        fmt.Println("getting error here only", err)
        return
     }

     defer c.Logout()

     service := c.Service

     chasis, _ := service.Chassis() 
     
     for _, chass := range chasis {
         fmt.Printf("Chasis: %#v\n\n", chass)
     }

}

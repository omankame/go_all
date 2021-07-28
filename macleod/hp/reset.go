package main

import (
        "fmt"
//        "net/http"
        "github.com/stmcginnis/gofish"
        "github.com/stmcginnis/gofish/redfish"
)

func main() {
     
     username := "hpadmin"
     config := gofish.ClientConfig {
               Endpoint: "https://172.16.2.84",
               Username: username,
               Password: "hpinvent",
               Insecure: true,
              }

     c, err := gofish.Connect(config)
     if err != nil {
        fmt.Println(err)
     }
     
     defer c.Logout()

     service := c.Service

     systems, err := service.Systems() // ([]*redfish.ComputerSystem, error) 
     if err != nil {
        panic(err)
     }

     for _, sys := range systems {
         err :=  sys.Reset(redfish.ForceRestartResetType)
         if err != nil {
            panic(err)
         }
     }
}

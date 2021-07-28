package main

import (
        "fmt"
//        "net/http"
        "github.com/stmcginnis/gofish"
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
     fmt.Println(c)
     defer c.Logout()

     service := c.Service

     systems, err := service.Systems() // ([]*redfish.ComputerSystem, error) 
     if err != nil {
        panic(err)
     }

     for _, sys := range systems {
          fmt.Printf("Systems Info: %#v\n\n", sys)
     }

 
     chassis, err := service.Chassis()
     if err != nil {
        panic(err)
     }

     for _, chass := range chassis {
		fmt.Printf("Chassis: %#v\n\n", chass)
	}


}


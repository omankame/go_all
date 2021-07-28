package main

import (
        "github.com/stmcginnis/gofish"
        "fmt"
//        "github.com/stmcginnis/gofish/redfish"
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

     ss, _ := service.Systems()
       
     
     for _, system := range ss {
           
           fmt.Printf("%T\n", system)
         na, _ :=  system.PCIeDevices() 
         for _, v := range na.pcieDevices {
             fmt.Println(v)
         }          
     }
//networkAdapters:"/redfish/v1/Chassis/1/NetiiworkAdapters/"
}

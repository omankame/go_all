package main

import (
        "github.com/stmcginnis/gofish"
        "fmt"
//        "os"
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

/*       memsum :=  redfish.MemorySummary {
             TotalSystemMemoryGiB: redfish.MemorySummary,
       }  */ 

/* get the value from "type MemorySummary struct " of link https://github.com/stmcginnis/gofish/blob/main/redfish/computersystem.go  

system is of type []*redfish.ComputerSystem

// ComputerSystem is used to represent resources that represent a // computing system in the Redfish specification. 
type ComputerSystem struct {  read this structure always as it conatins all the parameter to get data vimp
     MemorySummary MemorySummary
}  */


     
     for _, system := range ss {
            totalmemGB := system.MemorySummary.TotalSystemMemoryGiB
            fmt.Println("TOTAL RAM IS ", totalmemGB)    
     } 

    

}

package main

import (
       "fmt"
       "github.com/stmcginnis/gofish"
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
        panic(err)
     }

     defer c.Logout()

     service := c.Service

     ss, _ := service.Systems()
     for _, system := range ss {
          bootOrder := system.Boot.BootSourceOverrideMode
          fmt.Println(bootOrder)
     }

}

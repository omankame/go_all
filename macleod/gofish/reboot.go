package main

import (
	"fmt"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/redfish"
        "os"
)

func main() {
   
        args := os.Args[1:]
        if len(args) != 1 {
           fmt.Println("please enter https://ip-address only")
           return
        }
            
       
	// Create a new instance of gofish client, ignoring self-signed certs
	config := gofish.ClientConfig{
		Endpoint: args[0],
		Username: "hpadmin",
		Password: "5g_aci_36s74s5k1sc_no_os",
		Insecure: true,
	}

	c, err := gofish.Connect(config)
	if err != nil {
		panic(err)
	}
	defer c.Logout()

        service := c.Service

        ss, err := service.Systems()
	if err != nil {
		panic(err)
	}

	// Creates a boot override to pxe once
	bootOverride := redfish.Boot{
		BootSourceOverrideTarget:  redfish.HddBootSourceOverrideTarget,
		BootSourceOverrideEnabled: redfish.DisabledBootSourceOverrideEnabled,
	}

	for _, system := range ss {
		fmt.Printf("System: %#v\n\n", system)
		err := system.SetBoot(bootOverride)
		if err != nil {
			panic(err)
		}
		err = system.Reset(redfish.ForceRestartResetType)
		if err != nil {
			panic(err)
		}
	}
} 

        

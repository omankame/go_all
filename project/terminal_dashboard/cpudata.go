package main

import (
	"log"
        "fmt"
	linuxproc "github.com/c9s/goprocinfo/linux"
)

func main() {
stat, err := linuxproc.ReadStat("/proc/stat")
if err != nil {
	log.Fatal("stat read fail")
}

for _, s := range stat.CPUStats {
	 s.User
	 s.Nice
	 s.System
//	 s.Idle
//	 s.IOWait
      fmt.Println( s.User, "/t", s.Nice,  "/t", s.System)
}
}

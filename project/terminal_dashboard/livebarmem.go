package main

import (
        "fmt"
        "io/ioutil"
        "strings"
        "strconv"
        "time"
)

func main() {
    for {

     ToM := getTotalMem()
     fmt.Printf("Total mem = %v MiB\n", ToM)      
     time.Sleep(3 * time.Second)
   }
}
     

func getTotalMem() (Tom uint64) {
     contents, err := ioutil.ReadFile("/proc/meminfo")
     if err != nil {
     return
     }

     lines := strings.Split(string(contents), "\n")
//     fmt.Printf("%[1]v %[1]T", lines[0])
//     fmt.Printf("%v\n", lines[0])
     fields := strings.Fields(lines[0])
//     fmt.Printf("%v %v \n", fields[0], fields[1]) 
    
    umem, _ :=  strconv.ParseUint(fields[1], 10, 64) 
//  fmt.Printf("Total mem = %v MiB\n", bToMb(umem))
    fval:= bToMb(umem)
    return fval
}

func bToMb(b uint64) uint64 {
     return b / 1024 / 1024
}


     

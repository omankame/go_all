package main

import (
        "runtime"
        "fmt"
//        "time"
)

func main() {
     var m runtime.MemStats
     runtime.ReadMemStats(&m)
     
     mbV := bToMb(m.Alloc)
     fmt.Printf("Alloc = %v MiB\n", mbV)
     fmt.Printf("Alloc = %v MiB\n", bToMb(m.TotalAlloc))
     fmt.Printf("Alloc = %v MiB\n", bToMb(m.Sys))

}

func bToMb(b uint64) uint64 {
     return b / 1024 / 1024
}

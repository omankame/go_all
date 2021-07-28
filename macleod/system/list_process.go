package main

import (
        "fmt"
        "os/exec"
        "os"
        "syscall"
)

func main() {

     PS, err := exec.LookPath("ps")
     if err != nil {
        fmt.Println(err)
     }
     fmt.Println(PS)
     env := os.Environ()
//     fmt.Printf("%v\n", env)

    command := []string{ "PS", "-a", "-x"}

    err = syscall.Exec(PS, command, env) 
} 

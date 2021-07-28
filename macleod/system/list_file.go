package main

import (
        "os"
        "syscall"
        "os/exec"
        "fmt"
)

func main() {

      LS, _ := exec.LookPath("ls")
      fmt.Println(LS)

      command := []string{"ls", "-lrth",}
     
     env := os.Environ() 
     syscall.Exec(LS, command, env) 
}  


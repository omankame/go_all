// importatn package involved in system programming in Go ==> os, sysc.., flag, filepath

package main

import (
        "os"
        "path/filepath"
        "fmt"
        "log"
)

func main() {

    args := os.Args
    pwd, err := os.Getwd()
    if err != nil {
       fmt.Println("Error:", err)
    }

    if len(args) == 1 {
    return
    }

    if args[1] != "-P" {
    return
    }

    fileinfo, err := os.Lstat(pwd)
    if fileinfo.Mode()&os.ModeSymlink != 0 {
        realpath, err := filepath.EvalSymlinks(pwd)
        if err != nil {
            log.Println(err)
        }
        fmt.Println(realpath)
    }
}




package main

import (
        "strconv"
        "encoding/binary"
        "bytes"
        "fmt"
        "os"
)

func main() {
     args := os.Args[1:]
     if len(args) != 1 {
     fmt.Println("please enter the number to convert into binary")
     return
     }

     num, err  := strconv.ParseInt(args[0], 10, 64) 
     if err != nil {
     panic(err)
     os.Exit(1)
     }
     buf := new(bytes.Buffer)
     binary.Write(buf, binary.BigEndian, num) 
     fmt.Printf("%d is %x in BigIndian. \n", num, buf)

     buf.Reset()

     var b bytes.Buffer
     binary.Write(&b, binary.LittleEndian , num)
     fmt.Printf("%d is %x in LittleIndian. \n", num, b) 
}

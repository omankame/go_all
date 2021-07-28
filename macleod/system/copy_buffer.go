package main

import (
        "fmt"
        "os"
        "io"
        "strconv"
)

func main() {
     args := os.Args[1:]
     if len(args) != 3 {
     fmt.Println("usage: source destination BUFFERSIZE")
     os.Exit(1)
     }

     src := args[0]
     dst := args[1]
     buf, _ := strconv.Atoi(args[2])
     fmt.Println(src, dst, buf)
     err := copy(src, dst, buf)
     if err != nil {
       fmt.Println(err)
       fmt.Println("Error occured while copying file")
     }
    
}

func copy(source string, dest string, buffer int)  error {
      statSource, _ := os.Stat(source)
 
      if !statSource.Mode().IsRegular()  {
          fmt.Println("source file is not regular file")
          os.Exit(1)
      }

      sf, _ := os.Open(source)

      defer sf.Close()

      sbuf := make([]byte, buffer)
  
      _, err  := os.Stat(dest)
      if err == nil {
         fmt.Println("Destination File already exists, can not create and/or overwrite file")
         os.Exit(1)
      }

      df, _ := os.Create(dest)
      defer df.Close()

      for {
            n, err := sf.Read(sbuf) 
           if err != nil && err != io.EOF {
              return err
           }
           if n == 0 {
              break
           }

           _, err = df.Write(sbuf[:n]) 
           if err != nil {
              fmt.Println("error while writing data")
              return err
            }

         }
     return nil
}





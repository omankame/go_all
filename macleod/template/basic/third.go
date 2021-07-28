package main

import ( 
        "fmt"
        "os"
        "log"
        "io"
        "strings"
       )

func main() {
         args := os.Args[1:]
         if len(args) != 1 {
            log.Fatal("Please enter name ")
            os.Exit(1)
         }
         name := args[0]
         

str := fmt.Sprint(`
                <!DOCTYPE html>
                <html lang="en">
                <head>
                <meta charset="UTF-8">
                <title>Hello World!</title>
                </head>
                <body>
                <h1>` +
                name +
                `</h1>
                </body>
                </html>
        `)

       fmt.Println(str)

       nf, err := os.Create("index3.gohtml")
       if err != nil {
          log.Fatal(err)
       }

       defer nf.Close()

       io.Copy(nf, strings.NewReader(str))

}



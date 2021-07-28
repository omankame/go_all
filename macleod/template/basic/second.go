package main

import (
         "fmt"
         "os"
         "log"
         "io"
         "strings"
       )

func main() {
         name := "Todd Mcleod"

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

         
      nf, err := os.Create("index.html")
      if err != nil {
         log.Fatal(err)
         os.Exit(1)
       }
       defer nf.Close()

       io.Copy(nf, strings.NewReader(str))


}

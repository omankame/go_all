package main

import (
        "github.com/jung-kurt/gofpdf"
        "os"
        "fmt"
)

const str = `Please enter the name of image and description if required.
             go run second.go "image.jpg" "Aadhar Card" or
             go run second.go "image.jpg" ""
             `
func main() {
        args := os.Args[1:]
        if len(args) != 2 {
        fmt.Println(str)
        return
        }
        err := genratePdf("hello.pdf", args)
        if err != nil {
        panic(err)
        }

}


func genratePdf(filename string, args []string) error {
 
     pdf := gofpdf.New("P", "mm", "A3", "")
     pdf.AddPage()
     pdf.SetFont("Arial", "B", 16)
     pdf.CellFormat(190, 7, args[1], "0", 0, "CM", false, 0, "")
     pdf.ImageOptions(
                      args[0], 
                       50, 20,
                       150, 200,
                       false,
                       gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 
                       0,
                       "",
              ) 


     return pdf.OutputFileAndClose(filename) 

}        
         

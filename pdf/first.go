package main

import (
        "github.com/jung-kurt/gofpdf"
)

func main() {
        err := genratePdf("hello.pdf")
        if err != nil {
        panic(err)
        }

}

func genratePdf(filename string) error {
 
     pdf := gofpdf.New("P", "mm", "A4", "")
     pdf.AddPage()
     pdf.SetFont("Arial", "B", 16)
     pdf.CellFormat(190, 7, "Welcome to golangcode.com", "0", 0, "CM", false, 0, "")

     pdf.ImageOptions(
                      "golang-300x250.png", 
                       80, 20,
                       0, 0,
                       false,
                       gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 
                       0,
                       "",
              )

     return pdf.OutputFileAndClose(filename) 

}        
         

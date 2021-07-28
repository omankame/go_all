package main

import (
      "fmt"
//    "os"
      "github.com/360EntSecGroup-Skylar/excelize"
//      "html/template"
//      "net/http"
)




func main() {

    f, err := excelize.OpenFile("./Book99.xlsx")
    if err != nil {
        panic(err)
    }
    sheetName := f.GetSheetName(0)

    // Set cell A1 to Wrap
    style, err := f.NewStyle(&excelize.Style{
        Alignment: &excelize.Alignment{
            WrapText: true,
        },
    })
    if err != nil {
        panic(err)
    }
    if err := f.SetCellStyle("Sheet1", "A1", "A1", style); err != nil {
        panic(err)
    }

    // Get cell value, add new name and save
    v, err := f.GetCellValue(sheetName, "a1")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%x\n", v) // Output as hex so you can see how the line break is encoded
    err = f.SetCellValue(sheetName, "a1", v + "\n" + "Joe Bloggs")
    if err != nil {
        panic(err)
    }
    err = f.SaveAs("./Book99.xlsx")
    if err != nil {
        panic(err)
    }
}

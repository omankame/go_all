package main

import (
    "fmt"
//    "os"
    "github.com/360EntSecGroup-Skylar/excelize"
)

func main() { 
     

    f, err := excelize.OpenFile("./Book2.xlsx")
     if err != nil {
        fmt.Println(err)
        return
     }

    sheetname := f.GetSheetName(0)  // GetSheetName provides a function to get the sheet name of the workbook by the given sheet index
    fmt.Println(sheetname)   
    rows, err := f.GetRows(sheetname)

    rlast := len(rows)
 
    fmt.Println(rlast)

// https://dzone.com/articles/go-language-library-for-reading-and-writing-micros and
// https://stackoverflow.com/questions/28779596/find-last-filled-row-in-excel-using-golang


}

package main

import (
        "fmt"
//    "os"
    "github.com/360EntSecGroup-Skylar/excelize"
)

type M map[string]interface{}

var data = []M{
         M{"Name": "ONKAR", "GENDER": "MALE", "AGE": 31},
         M{"Name": "NIKHIL", "GENDER": "MALE", "AGE": 28},
         M{"Name": "UTTARA", "GENDER": "FEMALE", "AGE": 28},
       }

    


func main() {
     xlsx := excelize.NewFile()
     sheetName := "Sheet1"
     xlsx.SetSheetName(xlsx.GetSheetName(1), sheetName)

     xlsx.SetCellValue(sheetName, "A1", "Name")
     xlsx.SetCellValue(sheetName, "B1", "GENDER")
     xlsx.SetCellValue(sheetName, "C1", "AGE")

    for i, each := range data {
        xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), each["Name"])
        xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), each["GENDER"])
        xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), each["AGE"])
       }

    err := xlsx.SaveAs("./Book99.xlsx")
    if err != nil {
       fmt.Println(err)
    }

}

 
        
       

     





         


package main

import (
        "fmt"
    "os"
    "github.com/360EntSecGroup-Skylar/excelize"
    "strconv"
)

type M map[string]interface{}

/*var data = []M{
         M{"Name": "ONKAR", "GENDER": "MALE", "AGE": 31},
         M{"Name": "NIKHIL", "GENDER": "MALE", "AGE": 28},
         M{"Name": "UTTARA", "GENDER": "FEMALE", "AGE": 28},
       } */

    


func main() {
     args := os.Args[1:]
     if len(args) != 3 {
     fmt.Println("Please enter name , gender and age")
     return
     }
        
     age_num, err := strconv.Atoi(args[2])  
     var data = M{"Name": args[0], "GENDER": args[1], "AGE": age_num}
     
//     xlsx := excelize.NewFile()  ==> sheet creation part
//     sheetName := "Sheet1"
//     xlsx.SetSheetName(xlsx.GetSheetName(1), sheetName)


      f, err := excelize.OpenFile("./Book99.xlsx")  // ==> sheet open
      if err != nil {
      fmt.Println(err)
        return
     }

    sheetName := f.GetSheetName(0)
    rows, err := f.GetRows(sheetName)  // ==> last row of shhet
    rlast := len(rows)
    fmt.Println(rlast)



//     xlsx.SetCellValue(sheetName, "A1", "Name")
//     xlsx.SetCellValue(sheetName, "B1", "GENDER")
//     xlsx.SetCellValue(sheetName, "C1", "AGE")

//    for i_, each := range data {
        f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), data["Name"])
        f.SetCellValue(sheetName, fmt.Sprintf("B%d", rlast+1), data["GENDER"])
        f.SetCellValue(sheetName, fmt.Sprintf("C%d", rlast+1), data["AGE"])
//       }

    err = f.SaveAs("./Book99.xlsx")
    if err != nil {
       fmt.Println(err)
    }

}

 
        
       

     





         


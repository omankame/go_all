package main

import (
        "fmt"
        "io/ioutil"
        "encoding/json"
        "github.com/360EntSecGroup-Skylar/excelize"
        "sort"
)


func main() {

     file, err  := ioutil.ReadFile("test.json")
     if err != nil {
        panic(err)
     }

     var data interface{}
     json.Unmarshal(file, &data) // reading all json data

//     fmt.Println(data)

     t, ok := data.([]interface{})  // assertion in Interface
     _ = ok
      fmt.Printf("%[1]v  %[1]T\n", t)  // got the value of map[Age:28 Gender:female Name:Uttara]  map[string]interface {}

     myMap := t[0]   // aim here to get Age, Gender and Name which will be column name of Excel sheet

     columnData, _  := myMap.(map[string]interface {})  // extract the underlying concrete data from interface i.e type assertion

     fmt.Printf("%T\n", columnData)

     keys := make([]string, 0, len(columnData)) // creating and initializing slice to store column Name

     for k := range columnData {
         fmt.Printf("%[1]v %[1]T\n", k)
         keys = append(keys, k)
      }

      sort.Strings(keys)

      excelGenerate(keys)
}

func excelGenerate(keys []string) {

     xlsx := excelize.NewFile()
     sheetName := "Sheet1"
     xlsx.SetSheetName(xlsx.GetSheetName(1), sheetName)

     c := 'A'
    asciiValue := int(c)
    var a string
    for i := 0; i < len(keys); i++ {
        a = string(asciiValue)
        xlsx.SetCellValue(sheetName, a + "1", keys[i])
        asciiValue++
    }

    err := xlsx.SaveAs("./Onkar.xlsx")
    if err != nil {
       fmt.Println(err)
       return
    }

    fmt.Println("Excel file generated sucessfully")


}

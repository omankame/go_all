package main

import (
        "fmt"
        "io/ioutil"
        "encoding/json"
        "github.com/360EntSecGroup-Skylar/excelize"
)


func main() {

     filename := "test.json"
     var data interface{}
     plan, _ := ioutil.ReadFile(filename)

     json.Unmarshal(plan, &data)
 
     t, _ := data.([]interface{})
//     fmt.Printf("%T\n", t)
//     fmt.Println(t)

     f, err := excelize.OpenFile("./Book999.xlsx")
     if err != nil {
     fmt.Println(err)
     return
     }

    sheetName := f.GetSheetName(0)
    fmt.Println(sheetName)
    rows, err := f.GetRows(sheetName)
    rlast := len(rows)

    var concat string
 
     for i:= 0; i < len(t); i++ {
       rlast++
       myMap := t[i]
       final, _  := myMap.(map[string]interface{})
       values := make([]string, 0, len(final))
          for _, v := range final {
          finalValue, _ := v.(string)
          values = append(values, finalValue)
          }
       fmt.Println(values, len(values))
       fmt.Println(values[0])
       c := 'A'
       asciiValue := int(c)
       var a string
       for i := 0; i < len(values); i++ {
        a = string(asciiValue) 
        fmt.Println("I am in ascii", rlast)   
        concat = fmt.Sprint(a, rlast)     
        f.SetCellValue(sheetName, concat, values[i])
        asciiValue++
        }
      
     }

      
     err = f.SaveAs("./Book999.xlsx")
     if err != nil {
        fmt.Println(err)
     }





}


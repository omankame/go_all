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

/*     var m map[string]interface{} very important to extract the value https://stackoverflow.com/questions/41135686/how-can-we-read-a-json-file-as-json-object-in-golang  */

     t, ok := data.([]interface{})

     fmt.Println(t, ok)

     fmt.Printf("%v\n", t[0])

     myMap := t[0]
     final, _  := myMap.(map[string]interface{})
     keys := make([]string, 0, len(final))
    for k := range final {
       keys = append(keys, k)
    }

    fmt.Println(keys, len(keys))
    fmt.Println(keys[0])

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

    err := xlsx.SaveAs("./Book999.xlsx")
    if err != nil {
       fmt.Println(err)
    }



}

package main

import (
        "fmt"
        "io/ioutil"
        "encoding/json"
        "github.com/360EntSecGroup-Skylar/excelize"
        "sort"
        "reflect"
)


func main() {

     file, err  := ioutil.ReadFile("test.json")
     if err != nil {
        panic(err)
     }

     var data interface{}
     json.Unmarshal(file, &data) // reading all json data

     fmt.Println(data) 

     t, ok := data.([]interface{})  // assertion in Interface
     _ = ok
      fmt.Printf("%[1]v  %[1]T\n", t)  // got the value of map[Age:28 Gender:female Name:Uttara]  map[string]interface {}


     

     myMap := t[0]   // aim here to get Age, Gender and Name which will be column name of Excel sheet

     columnData, _  := myMap.(map[string]interface {})  // extract the underlying concrete data from interface i.e type assertion
  
     fmt.Printf("Check ----------------%T\n", columnData)

//     keys := make([]string, 0, len(columnData)) // creating and initializing slice to store column Name
     keys := Sorting(columnData)
//     for k := range columnData {
//         fmt.Printf("%[1]v %[1]T\n", k)
//         keys = append(keys, k)
//      }
      
     sort.Strings(keys)
     fmt.Println(keys)
     excelGenerate(keys)
//     openExcelAndPrintColumn()
      openExcelAndInsertData(data)
}

func Sorting(columnData map[string]interface {}) []string {
    keys := make([]string, 0, len(columnData)) 
    for k := range columnData {
         fmt.Printf("%[1]v %[1]T\n", k)
         keys = append(keys, k)
      }
      
     return keys
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

/* func openExcelAndPrintColumn() {

     f, err := excelize.OpenFile("./Onkar.xlsx")
     if err != nil {
        fmt.Println(err)
        return
     }
     
     sheetName := f.GetSheetName(1)

     cols, _ := f.GetCols(sheetName) 
     fmt.Println(cols)

    
    rows, err := f.GetRows(sheetName)
    rlast := len(rows)
    fmt.Println(rlast)

    

} */


func openExcelAndInsertData(data interface{} ) {

     f, err := excelize.OpenFile("./Onkar.xlsx")            // to open the excel shhet
     if err != nil {
        fmt.Println(err)
        return
     }


    sheetName := f.GetSheetName(0)           // getting the sheet name
    rows, err := f.GetRows(sheetName)        // get total rows
    rlast := len(rows)                       // calculate length reuired in future to put next data in which row
    fmt.Println(rlast)
    fmt.Println(sheetName)


     t, _ := data.([]interface{})    // data hold as interface, then with type assertion stored in t
      mData :=  t[0].(map[string]interface {})
     
       keys := Sorting(mData)
       sort.Strings(keys)
       fmt.Println(keys)
       for  i := 0; i < len(t); i++ {                 // in json lets say 5 data set exists then len(t) = 5
        mData, _  = t[i].(map[string]interface {})  // again with help of type assertion extracted data in mData

        
//        keys := make([]string, 0, len(mData)) // creating and initializing slice to store column Name  
//          for k := range mData {
//            fmt.Printf("%[1]v %[1]T\n", k)
//            keys = append(keys, k)
//          }

//        sort.Strings(keys)
        test := reflect.ValueOf(mData).MapKeys()
         c := 'A'                                  // same logic use in above function to get the column name
           asciiValue := int(c)
           var a string
             for j := 0; j < len(keys) ; j++  {     // another for loop to get the data and put in proper cell of Excel sheet
                 fmt.Println(mData[ test[j].Interface().(string) ] ) 
                 a = string(asciiValue)
                exData := mData[ test[j].Interface().(string) ]
                f.SetCellValue(sheetName, a + fmt.Sprintf("%d", rlast +1) , exData ) 
 
                asciiValue++                                                  // become 'B'

//                f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), mData[ keys[j].Interface().(string) ])
            }
            rlast++  // for third, fourth , fifth ... row
       }

 err = f.SaveAs("./Onkar.xlsx")   // finally save the data
     if err != nil {
        fmt.Println(err)
     }

   fmt.Println("JSON TO EXCEL CONVERSION DONE SUCESSFULLY")
                                                               
}         


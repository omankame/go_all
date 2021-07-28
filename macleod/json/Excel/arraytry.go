// sequence sathi modification lagel
//https://stackoverflow.com/questions/18342784/how-to-iterate-through-a-map-in-golang-in-order
//https://bitfieldconsulting.com/golang/map-string-interface


package main

import (
        "fmt"
        "encoding/json"
        "io/ioutil"
        "reflect"
        "github.com/360EntSecGroup-Skylar/excelize"
//        "sort"
)


func main() {
     fileByte, err  := ioutil.ReadFile("test.json")  // to read the json file as byte
     if err != nil {
        panic(err)
     }

     var data  interface{}     // store data 

     json.Unmarshal(fileByte, &data)  //json unmarshalling, required address &data is important

     colNameExtract(data)    // calling function to get first json data, then extract column value , generate excel and save it.

     openExcelAndInsertData(data)   // last function, open the excel again and extract the data with help of reflect and put in proper cell
}

func colNameExtract(data interface{} ) {

     t, _ := data.([]interface{})    // extracte concrete data under the interface using type assertion
     fmt.Printf("%T \n", t)         // getting the type of t
 
     columnData, _  := t[0].(map[string]interface {})    // type assertion with first instance of json file
     fmt.Printf("%[1]T %[1]v \n", columnData)  //map[string]interface {} map[Age:28 Gender:female Name:Uttara] checking its type and actual da

     
     
//    keys := reflect.ValueOf(columnData).MapKeys()  // with the help of reflection extracted the keys from Map which will be coulmn name of xl
      val := reflect.ValueOf(columnData)

      keys := val.MapKeys()
      for _, k := range keys {
           c_key := k.Convert(val.Type().Key())        
           fmt.Printf("%[1]T %[1]v \n", c_key)
      }
    

//sorting of keys as map always generate random sequence, at least to keep sequnce in some meaningful way
    fmt.Printf("%[1]T %[1]v \n", keys)              // checking the value of keys and its type which is []reflect.Value
    fmt.Println(keys[0], keys[1], keys[2] , len(keys))   
      
 
    excelGenerate(keys)     // call one more function to generate excel sheet
} 

func excelGenerate(keys []reflect.Value) {   

     xlsx := excelize.NewFile()                           // to create new file
     sheetName := "Sheet1"                                // to create new sheet in xl file
     xlsx.SetSheetName(xlsx.GetSheetName(1), sheetName)   

     c := 'A'                                 // to set the column name starting from "A1" of excel
     asciiValue := int(c)
     var a string

     for i := 0; i < len(keys) ; i++ {             // keys nothing but column name in excel sheet 
         a = string(asciiValue)                    
         xlsx.SetCellValue(sheetName, a + "1", keys[i])   // A1 = Name, B1 = Gender,  C1 = Age like that 
         asciiValue++                                     // A becomes B
     }

     err := xlsx.SaveAs("./OM.xlsx")             // Save the file
     if err != nil {
       fmt.Println(err)
       return
     }

}        

func openExcelAndInsertData(data interface{} ) {


     f, err := excelize.OpenFile("./OM.xlsx")            // to open the excel shhet
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

      for  i := 0; i < len(t); i++ {                 // in json lets say 5 data set exists then len(t) = 5
       	mData, _  := t[i].(map[string]interface {})  // again with help of type assertion extracted data in mData 
       	keys := reflect.ValueOf(mData).MapKeys()     // vimp to get the column name as keys. keys[0] = Name
//        sort.Strings(keys)
           c := 'A'                                  // same logic use in above function to get the column name  
           asciiValue := int(c)       
           var a string
             for j := 0; j < len(keys) ; j++  {     // another for loop to get the data and put in proper cell of Excel sheet
                
                fmt.Println(mData[ keys[j].Interface().(string) ] )   // getting actual Data with help of reflection.Interface() function
                                                      // https://stackoverflow.com/questions/17262238/how-to-cast-reflect-value-to-its-type

                a = string(asciiValue)
                exData := mData[ keys[j].Interface().(string) ]
                f.SetCellValue(sheetName, a + fmt.Sprintf("%d", rlast +1) , exData )  // actual data put in excel sheet, 
                                                                               // ex. a + fmt.Sprintf("%d", rlast +1) ==> 'A2' 
                asciiValue++                                                  // become 'B'

//                f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), mData[ keys[j].Interface().(string) ])
            }
             rlast++  // for third, fourth , fifth ... row 
       }              

     err = f.SaveAs("./OM.xlsx")   // finally save the data
     if err != nil {
        fmt.Println(err)
     }

   fmt.Println("JSON TO EXCEL CONVERSION DONE SUCESSFULLY")

}

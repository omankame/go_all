package main

import (
        "fmt"
//    "os"
    "github.com/360EntSecGroup-Skylar/excelize"
    "html/template"
    "net/http"
    "strings"
)

var tpl *template.Template
type M map[string]interface{}

func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
     http.HandleFunc("/", datasubmit)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func datasubmit(w http.ResponseWriter, req *http.Request) {

      nm := req.FormValue("name")
      ge := req.FormValue("gender")
      ag := req.FormValue("age")

      sliNm := strings.Split(nm, ";")
       var final string
      for _, v := range sliNm {

      final += v 
      final = final  + "\n"
    }

    fmt.Println(final)


      
      var data = M{"Name": final, "GENDER": ge, "AGE": ag}

      tpl.ExecuteTemplate(w, "index.gohtml", data)

      f, err := excelize.OpenFile("./Book99.xlsx")
          if err != nil {
          fmt.Println(err)
          return
         }

      sheetName := f.GetSheetName(0)
       // Set cell A1 to Wrapp
      
     style, err := f.NewStyle(&excelize.Style{
          Alignment: &excelize.Alignment {
             WrapText: true,
          },    
         })  
     if err != nil {
            panic(err)
     }
  
//     if err := f.SetCellStyle("Sheet1", "A1", "A1", style); err != nil {
//        panic(err)
//     }


//    fmt.Printf("%x\n", v)
    

      fmt.Println(sheetName)
      rows, err := f.GetRows(sheetName)
      rlast := len(rows)
      fmt.Println(rows)
      fmt.Println(rlast)
 
            
      f.SetCellStyle(sheetName, fmt.Sprintf("A%d", rlast+1), fmt.Sprintf("A%d", rlast+1), style);
      fmt.Println(data["Name"]) 
      f.SetCellValue(sheetName, fmt.Sprintf("A%d", rlast+1), data["Name"] )
      f.SetCellValue(sheetName, fmt.Sprintf("B%d", rlast+1), data["GENDER"])
      f.SetCellValue(sheetName, fmt.Sprintf("C%d", rlast+1), data["AGE"])
 
      err = f.SaveAs("./Book99.xlsx")
       if err != nil {
        fmt.Println(err)
       }


}

 
        
       

     





         


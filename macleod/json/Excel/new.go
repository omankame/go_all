/* https://stackoverflow.com/questions/36752050/converting-nested-json-to-csv-in-go
https://blog.golang.org/json
https://stackoverflow.com/questions/52433236/how-to-access-key-and-value-from-json-array-in-go-lang
https://stackoverflow.com/questions/52433236/how-to-access-key-and-value-from-json-array-in-go-lang */

package main

import (
        "fmt"
        "io/ioutil"
        "strconv"
        "encoding/json"
//        "sort"
//        "reflect"
//        "github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
            file, err  := ioutil.ReadFile("test.json")
                  if err != nil {
                                  panic(err)
                  }

            var data interface{}
            err = json.Unmarshal(file, &data) // reading all json data
            if err != nil {
               fmt.Println(err)
            }

            
            t, _ := data.([]interface{}) 
            fmt.Printf("%[1]v %[1]T \n", t)

            for  i := 0; i < len(t); i++ {  
//              for _, m := range t.(map[string]interface {}) {
                m, _  := t[i].(map[string]interface {})
               decodeJson(m)
//              fmt.Printf("%T \n", m) //its type map[string]interface {}
           }

}

  func decodeJson(m map[string]interface {}) []string {

      values := make([]string, 0, len(m))
//      keys   := make([]string, 0, len(m))
    for _, v := range m {
//        fmt.Println(k)
        switch vv := v.(type) {
        case map[string]interface{}:
            for _, value := range decodeJson(vv) {
                values = append(values, value)
            }
        case string:
            values = append(values, vv)
        case float64:
            values = append(values, strconv.FormatFloat(vv, 'f', -1, 64))
        case []interface{}:
            // Arrays aren't currently handled, since you haven't indicated that we should
            // and it's non-trivial to do so.
        case bool:
            values = append(values, strconv.FormatBool(vv))
        case nil:
            values = append(values, "nil")
        }
    }
    fmt.Println(values)
    return values
}  
 

//declare multidimensional array and then will try to iterate with respect key and then appending the value 

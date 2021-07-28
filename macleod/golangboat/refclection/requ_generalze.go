package main

import (
         "fmt"
         "reflect"
)

type order struct {
     ordId  int
     custId int
}


type employee struct {
     emp_name     string
     emp_id       int
     dept_id	  int
}
func main() {
     o := order{
          ordId: 100,
          custId: 99,
        }

     e := employee{
          emp_name: "onkar",
          emp_id:   99,
          dept_id:  9,
      }

     createQuery(o)

     createQuery(e)

}

func createQuery(q interface{} )  {
     t := reflect.TypeOf(q)
     v := reflect.ValueOf(q)
     if v.Kind() == reflect.Struct {
        query := fmt.Sprintf("Insert into %s values(", t.Name())
        for i:=0; i < v.NumField(); i++ {
            switch v.Field(i).Kind() {
                      case reflect.Int:
                if i == 0 {
                    query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
                } else {
                    query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
                }
            case reflect.String:
                if i == 0 {
                    query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
                } else {
                    query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
                }
            default:
                fmt.Println("Unsupported type")
                return
            }
       
         }
      query = fmt.Sprintf("%s)", query)
      fmt.Println(query)
      return
 }
fmt.Println("unsupported type")
} 
    

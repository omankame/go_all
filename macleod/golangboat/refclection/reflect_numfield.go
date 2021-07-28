package main

import (
        "fmt"
        "reflect"
)
type order struct {
     ordId     int
     custId    int
}

func main() {
     o := order{
          ordId: 100,
          custId: 99,
     }

     createQuery(o)

}

func createQuery(q interface{}) {
     t := reflect.TypeOf(q)
     v := reflect.ValueOf(q)
     if v.Kind() == reflect.Struct {
        fmt.Println("Total number of fields:", v.NumField())
       for i := 0; i < v.NumField(); i++ {
          fmt.Printf("Type is %T, Value is %v, Field number %d \n", t, v.Field(i), i)
          }
     }
}  

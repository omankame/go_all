package main

import (
         "fmt"
         "reflect"
)

type order struct {
     ordId  int
     custId int
}

func main() {
      o := order {
           ordId: 100,
           custId: 9,
        }

     createQuery(o)
//     fmt.Println(s)
}

func createQuery(q interface{})  {

    t := reflect.TypeOf(q)
    v := reflect.ValueOf(q)
    k := t.Kind()
    fmt.Println("Type ", t)
    fmt.Println("Value ", v)
    fmt.Println("Kind ", k)
   

//   i :=  fmt.Sprintf("Insert into Order values(%d, %d)", o.ordId, o.custId)
  
}

    

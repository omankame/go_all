package main

import (
         "fmt"
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

     s := createQuery(o)
     fmt.Println(s)
}

func createQuery(o order) string {
   i :=  fmt.Sprintf("Insert into Order values(%d, %d)", o.ordId, o.custId)
   return i
}

    

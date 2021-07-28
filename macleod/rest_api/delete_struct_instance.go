package main

import (
        "fmt"
)

type obj struct {
     a   int
     b   string
}

func main() {
     o1 := &obj{1, "onkar",}
     o2 := &obj{2, "nikhil",}

     fmt.Printf("%v\n", o1)
     fmt.Printf("%v\n", o2)

     m := map[int]**obj {
          1: &o1,
          2: &o2,
      }

      // fmt.Printf("%t %v \n", o1, o1)
       fmt.Println(m, m[1])
       o1 = nil
       fmt.Println(m, m[1])      
}
 

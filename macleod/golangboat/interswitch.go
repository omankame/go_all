package main

import (
        "fmt"
)

func main() {
 
     findType("onkar")
     findType(100)
     findType(99.99)      
}

func findType(i interface{}) {
     switch i.(type) {

      case string: 
           fmt.Printf("value is %s\n", i.(string))
  
      case int: 
           fmt.Printf("value is %d\n", i.(int))
   
      default:
       fmt.Println("Unknown Type")
    }

}

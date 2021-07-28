/* only last parameter of fun can be variadic. Variadic fun. internally convert arguments as slice. append func uses variadic func in background. normal slice passes to variadic func generates error. with syntactic sugar you can use variadic func. */

package main

import (
        "fmt"
)

func main() {

     find(89, 89,30,35,40)
}

func find(num int, nums ...int) {
     found := false
     for i, v := range nums {
         if num == v {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
         }
    }
 
    if !found {
       fmt.Println(num , "did not find.")
    }
}

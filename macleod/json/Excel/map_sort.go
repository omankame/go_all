package main

import (
         "fmt"
         "sort"
)


func main() {

     map1 := make(map[string]int)

     map1["Mon"]=1
     map1["Tue"]=2
     map1["Wed"]=3
     map1["Thu"]=4
     map1["Fri"]=5
     map1["Sat"]=6
     map1["Sun"]=7
     

    fmt.Println("UNsorted")

    for key, val := range map1 {
        fmt.Printf("%s is the %d of a week\n",key,val)
    }


    fmt.Println("Sorted")

    value := make([]int, len(map1))
    i := 0
    for _, val := range map1 {
        value[i] = val
        i++
    }

    sort.Ints(value)
    for i, _ := range value {
        for j := range map1 {
          if map1[j] == value[i] {
             fmt.Printf("%s is the %d of a week\n",j,value[i])
          }

        }
     }

}

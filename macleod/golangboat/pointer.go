package main

import (
         "fmt"
)

func main() {
    a := 58
    fmt.Println(a)

    b := &a
    change(b)
    fmt.Println("a after change", a)
} 

func change(val *int) {
     *val = 100
}

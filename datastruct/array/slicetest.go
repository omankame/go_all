package main

import (
         "fmt"
)

func main() {
     a := []int{55, 8, 90, 33,45, 22}
     fmt.Println(len(a) / 2)
     fmt.Println(a[:3])
     b := a[: len(a) / 2]
     fmt.Println(b)


     c := a[len(a) /2: ]
     fmt.Println(c)
     fmt.Println(a[5:])
}

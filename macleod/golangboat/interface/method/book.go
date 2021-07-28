package main

import (
         "fmt"
)

type book struct {
     title string
     price float64
}

func (b book) print() {
     fmt.Printf("Book Title %s \t Book Price %0.1f\n", b.title, b.price)
}


package main

import (
         "fmt"
)

type game struct {
     title string
     price float64
}

func (g game) print() {
     fmt.Printf("Game Title %s \t Game Price %0.1f\n", g.title, g.price)
}


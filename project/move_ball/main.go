package main

import (
        "fmt"
        "time"
        "math/rand"
)

const( 
      width = 76
      height = 10
      numBalls =1 
)

type screen[width][height]byte

type ball struct {
     x, y int
     xd, yd int
}

func main() {
     balls := []*ball{}
     
}

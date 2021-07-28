package main

import (
         "fmt"
         "errors"
         "math"
)
//const PI float64 = 3.14

func main() {
     radius := 20.5
     area, err := calArea(radius)
     if err != nil {
        fmt.Println(err)
        return 
     }

     fmt.Printf("Area of the circle is %0.2f\n", area)

}

func calArea(r float64) (float64, error) {
     if r < 0 {
        return 0, errors.New("The radius of the circle can not be negative")
     }

     return  math.Pi * r * r,  nil
} 

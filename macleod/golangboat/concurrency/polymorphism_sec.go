package main

import (
         "fmt"
)

type Income interface {
     calculate() int
     source() string
}

type FixedBillings struct {
     projectName   string
     biddedAmount  int
}

type TimeAndMaterials struct {
     projectName   string
     noOfHours	   int
     rateHour	   int
}

func (fb FixedBillings) calculate() int {
     return fb.biddedAmount
}

func (fb FixedBillings) source() string {
     return fb.projectName
}


func (tm TimeAndMaterials) calculate() int {
     return tm.noOfHours * tm.rateHour        
}

func (tm TimeAndMaterials) source() string {
     return tm.projectName
}

func main() {
     fib := FixedBillings{
            projectName: "JIO IMPLEMENTATION",
            biddedAmount: 100000,
     }

     tam := TimeAndMaterials{
            projectName: "JIO ANTENA",
            noOfHours:  100,
            rateHour:   20,    
     }

    finalIncome := []Income{fib, tam}
    finalCalculate(finalIncome)

}

func  finalCalculate(fi []Income) {
 
      totalVal := 0
      
      for _, v := range fi {
        fmt.Printf("The first Project name %s \n", v.source() )
        fmt.Printf("The first Project income %d$ \n", v.calculate() )
        
        totalVal = totalVal + v.calculate() 
       }

      fmt.Printf("Total income by different sources %d$ \n", totalVal )

}
           

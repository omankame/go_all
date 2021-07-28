package main

import (
        "fmt"
)

type employee interface {
     develop() int
     name() string
}

type team1 struct {
     totalapp_1 int
     name_1     string
}

type team2 struct {
     totalapp_2   int
     name_2       string
}

func (t1 team1) develop() int {
     return t1.totalapp_1
}

func (t1 team1) name() string {
     return t1.name_1
}

func (t2 team2) develop() int {
     return t2.totalapp_2
}

func (t2 team2) name() string {
     return t2. name_2
}

func main() {
     res1 := team1{
             totalapp_1: 20,
             name_1: "ANDROID",
     }

     res2 := team2{
             totalapp_2: 15,
             name_2: "IOS",
     }

     final := []employee{res1, res2}

     developTotal(final)

}

func developTotal(final []employee) {

     devTotal := 0
     for _, fi := range final {
         fmt.Printf("The Project environment %s\n", fi.name())
         fmt.Printf("Total number of project %d\n", fi.develop())
         devTotal = devTotal + fi.develop()
         }
    
     fmt.Printf("Total number of projects completed by Company: %d\n", devTotal)

}
         

 


// here we can add third team structe easily with its own method definationa and will calculate total number of project by COMPANY

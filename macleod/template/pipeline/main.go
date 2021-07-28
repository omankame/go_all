package main

import ( 
        "os"
        "text/template"
        "log"
        "math"
       )
var tpl *template.Template

var fm = template.FuncMap {
    "fdbl":   double,
    "fsq":    sqare,
    "fsqrt":  sqrt,
}

func init() {
     tpl = template.Must(template.New("cal.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func double(num int) int {
     return num + num
}

func sqare(num int) float64 {
    a :=  math.Pow(float64(num), 2)
    return a
}

func sqrt(nu float64) float64 {
     return math.Sqrt(nu)
}

func main() {

    err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 4)
    if err != nil {
       log.Fatal(err)
       os.Exit(1)
   } 

}


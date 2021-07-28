package main

import ( "fmt" )

func main() {
     a := []int{5, 4, 6, 99, 23, 1, 9}

     i := 1
     for i < len(a) {
          j := i
         for j >=1 {
             if a[j] < a[j-1] {
                a[j], a[j-1] = a[j-1], a[j]
             }
            j--
         }
        i++
     }

   fmt.Println(a)
}

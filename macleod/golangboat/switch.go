package main

import (
       "fmt"
)

func main()  {
     finger := 4
     fmt.Printf("Finger %d  is ", finger)
     switch  finger {
     case 1:
         fmt.Println("Thumb")
     case 2:
         fmt.Println("Index")
     case 3:
         fmt.Println("Middle")
     case 4:
         fmt.Println("Ringer")
     case 5:
         fmt.Println("Pinky")

    }


   letter := "i"

   fmt.Printf("The %s is ", letter)
   switch letter {
   case "a", "e", "i", "o", "u":
         fmt.Println("Vowels")
   default :
     fmt.Println("Consonant")
   }

}

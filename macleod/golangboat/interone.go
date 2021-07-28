package main

import (
        "fmt"
)
type VowelFinder interface {
     FindVowel() []rune 
}

type MyString string

func (ms MyString) FindVowel() []rune {
     var vowls []rune

     for _, v := range ms {
         if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
            vowls = append(vowls, v)
          }
    }
    return vowls
}
     


func main() {
      name := MyString("onkar mankame")
      var v VowelFinder   // declare interface as v
      v = name            // name can be assigned to v as name is of type MyString, Mystring implements interface VowelFinder
      fmt.Printf("Vowels are %c", v.FindVowel())

}

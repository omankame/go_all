package main

import (
         "fmt"
)

func main() {

     a := []int{66, 3, 4, 99, 22, 12, 5 , 90, 77, 123, 21}

   final :=   mergeSort(a)
   fmt.Println(final)
}

func mergeSort(n []int) []int {

     if len(n) == 1 {
        return n
     }

     var fp = mergeSort( n[: len(n) / 2 ] )
     var sp = mergeSort( n[len(n) / 2 :]  )
     
     return merge(fp, sp)

}

func merge(fp []int, sp []int) []int {
     var mArr = make([]int, len(fp) + len(sp))

     var fpIndex = 0 // I value
     var spIndex = 0 // J value
     var mArrIndex = 0 // K value

     for fpIndex < len(fp) && spIndex < len(sp) {
          if (fp[fpIndex] <= sp [spIndex]) {
             mArr[mArrIndex] = fp[fpIndex]
             fpIndex++
          } else {
            mArr[mArrIndex] = sp[spIndex]
            spIndex++
          }
          mArrIndex++
     }

    for fpIndex < len(fp) {
        mArr[mArrIndex] = fp[fpIndex]
        fpIndex++
        mArrIndex++
    }

    for spIndex < len(sp) {
        mArr[mArrIndex] = sp[spIndex]
        spIndex++
        mArrIndex++
    }


    return mArr

}

 
  

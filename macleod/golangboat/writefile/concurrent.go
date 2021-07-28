package main

import (
        "fmt"
        "os"
        "math/rand"
        "sync"
)

func main() {
     var wg sync.WaitGroup
    
     rchan := make(chan int)
     done := make(chan bool)
     for i := 0 ; i < 100; i++ {
        wg.Add(1)
        go produce(rchan, &wg)
     }
 
     go consume(rchan, done)

     wg.Wait()
     close(rchan)
     d := <-done
     if d == true {
         fmt.Println("All data written sucessfully")
     } else {
       fmt.Println("Failed writing file")
     }


}

func consume(rchan chan int, done chan bool) {
     f, err := os.Create("concurrent")
     defer f.Close()
     
     if err != nil {
        fmt.Println(err)
        return
     }

     for  v := range rchan {
        _, err := fmt.Fprintln(f, v)
           if err != nil {
              fmt.Println(err)
              done <- false
           }
      }
      done <- true
}      

func produce(rchan chan int, wg *sync.WaitGroup) {
     v := rand.Intn(999)

     rchan <- v
     wg.Done()
}

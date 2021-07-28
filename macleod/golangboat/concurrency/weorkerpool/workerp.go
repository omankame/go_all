package main

import (
        "fmt"
        "time"
)

func main() {

     // two channels to send them jobs and collect results
     jobs := make(chan int, 100)
     results := make(chan int, 100)
     
     // start up with 3 workers 
     for w := 1; w <= 3; w++ {
         go worker(w, jobs, results)
     }

     // send teh 5 jobs to and close, it indicates thats all the work we have done
     for j := 1; j <= 5; j++ {
         jobs <- j
     }

     close(jobs)

     // collect all the jobs of the work
     for a := 1; a <=5; a++ {
         <-results
     }

}      


func worker(id int, jobs chan int, results chan int) {
     for j := range jobs {
         fmt.Println("worker", id, "started job", j)
         time.Sleep(5 * time.Second)
         fmt.Println("worker", id, "finished job", j)
         results <- j * 2
      }
}

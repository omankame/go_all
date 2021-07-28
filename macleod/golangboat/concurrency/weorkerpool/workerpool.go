package main

import (
        "fmt"
        "time"
        "sync"
        "math/rand"
)

type Job struct {
     id    int
     randomno	int
}

type Result struct {
     job    Job
     sumofdigits  int
}

var jobs = make(chan int, 10)
var results = make(chan int, 10)

func main() {
     
     startTime := time.Now()
     noOfJobs := 100
     go allocate(noOfJobs)
     done := make(chan bool)
     go result(done)
     noofWorkers := 10
     createWorkerPool(noofWorkers)
     <-done
     endTime := time.Now()
     diff := endTime.Sub(startTime)
     fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

func allocate(noOfJobs int) {
     for i := 0; i < noOfJobs; i++ {
         random := rand.Intn(999)
         job := Job{i, random}
         jobs <- job
     }
     close(jobs)
}

func result(done chan bool) {
     for result := range results { 
         fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
     }
     done <- true
}

func createWorkerPool(noofWorkers int) {
     var wg sync.WaitGroup
     for i := 0; i < noofWorkers; i++ {
         wg.Add(1)
         go worker(&wg)
     }
    wg.Wait()

    close(results)
}

func worker(wg *sync.WaitGroup) {
     for j := range jobs {
         output := Result{j, digits(j.randomno)}
         results <- output
     }
    wg.Done()
}

func digits(number int) {
     sum := 0
    no := number
    for no != 0 {
        digit := no % 10
        sum += digit
        no /= 10
    }
    time.Sleep(2 * time.Second)
    return sum
}

       
     

 



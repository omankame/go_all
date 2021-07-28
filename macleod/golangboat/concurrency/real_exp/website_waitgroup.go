package main

import (
        "fmt"
        "sync"
        "net/http" 
)

var wg sync.WaitGroup


func main() {
    
     websites := []string{
                "https://stackoverflow.com/",
		"https://github.com/",
		"https://www.linkedin.com/",
		"http://medium.com/",
		"https://golang.org/",
		"https://www.udemy.com/",
		"https://www.coursera.org/",
		"https://wesionary.team/",
	}

    for _, v := range websites {
        wg.Add(1)
        go status(v, &wg)
    }
                         
    wg.Wait()
    fmt.Println("All go routines finished execution")

}

func status(web string, wg *sync.WaitGroup) {
     
     resp, err := http.Get(web)
     if err != nil {
        fmt.Println(err)
        return
     }
     if resp.StatusCode == 200 {
        fmt.Println(web, "IS Live")
        }
   
     wg.Done()
}

package main
import (
        "fmt"
        "net/http"
)


func main() {
     http.HandleFunc("/", homepage)
     http.ListenAndServe(":8080", nil)
}

func homepage(w http.ResponseWriter, req *http.Request) {
     fmt.Fprint(w, "Hello World!")
     fmt.Println("Endpoint Hit: HOME PAGE")
}

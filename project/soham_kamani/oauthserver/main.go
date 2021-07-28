package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/oauth2.v3/models"
	"log"
	"net/http"
	"time"

	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func main() {

     http.HandleFunc("/protected", func(w http.ResponseWriter, req *http.Request) {
                     w.Write([]byte("Hello, I am protected"))
        })

     http.ListenAndServe(":8080", nil)

}

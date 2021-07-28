package main

import (
        "fmt",
)

type song struct {
     name 	string
     artist	string
     next	*song
}

type playlist struct {
     name 	string
     head	*song
     playing	*song
}

func main() {

    play

}


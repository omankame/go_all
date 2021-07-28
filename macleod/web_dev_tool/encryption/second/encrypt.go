package main

import (
          "crypto/aes"
          "crypto/cipher"
          "crypto/rand"
          "fmt"
          "io"
         "io/ioutil"
          "os"
)

func main() {

     fmt.Println("Encryption Program")
     args := os.Args[1:]
     if len(args)  != 1 {
        fmt.Println("Please enter as string ")
        return
     }

     text := []byte(args[0])

     fmt.Println(text)

    key := []byte("passphrasewhichneedstobe32bytes!")

    c, err := aes.NewCipher(key)
    if err != nil {
       fmt.Println(err)
    }


    gcm, err := cipher.NewGCM(c)
    if err != nil {
       fmt.Println(err)
    }


    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
       fmt.Println(err)
    }

    fmt.Println(gcm.Seal(nonce, nonce, text, nil))

    err = ioutil.WriteFile("myfile.data", gcm.Seal(nonce, nonce, text, nil), 0777)
    if err != nil {
       fmt.Println(err)
    }

}
     

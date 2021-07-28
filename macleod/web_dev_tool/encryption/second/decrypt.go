package main

import (
        "io/ioutil"
        "crypto/aes"
        "crypto/cipher"
        "fmt"
)

func main() {
     fmt.Println("Decryption Program v1")

     key := []byte("passphrasewhichneedstobe32bytes!")
     ciphertext, err := ioutil.ReadFile("myfile.data")
     if err != nil {
        fmt.Println(err)
     }

     c, err := aes.NewCipher(key) 
     if err != nil {
        fmt.Println(err)
     }

     gcm, err := cipher.NewGCM(c)
     if err != nil {
        fmt.Println(err)
     }

     noncesize := gcm.NonceSize()
 
     if len(ciphertext) < noncesize {
        fmt.Println(err)
     }

     nonce, ciphertext := ciphertext[:noncesize], ciphertext[noncesize:]
     plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
     if err != nil {
        fmt.Println(err)
     }

     fmt.Println(string(plaintext))
}




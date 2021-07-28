package main

import (
         "crypto/aes"
         "crypto/cipher"
         "crypto/rand"
         "encoding/base64"
         "fmt"
         "io"
)

func main() {
     secret, err := Secret()
     if err != nil {
        panic(err)
     }

     fmt.Println(string(secret))

     message := "Very Top Information."
     fmt.Println("Message:", message)
   
     encrypted, err := encrypt(message, secret) 
     fmt.Println(encrypted)
}

// return 32 bytes as AES key
func Secret() ( []byte, error ) {
     key := make([]byte, 16)
     _, err := rand.Read(key)
     if err != nil {
        return nil, err
     }

     return key, nil


}

func encrypt(m string, s []byte) (string, error) {
     cipherB, err := aes.NewCipher(s)
     if err != nil {
        return "", err 
     }

     aead, err := cipher.NewGCM(cipherB)
     if err != nil {
        return "", err
     }

     nonce := make([]byte, aead.NonceSize())
     if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
           return "", err
     }

     return base64.URLEncoding.EncodeToString(aead.Seal(nonce, nonce, []byte(m), nil)), nil
}     



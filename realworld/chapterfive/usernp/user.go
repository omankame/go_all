package usernp

import (
        "golang.org/x/crypto/bcrypt"
        "fmt"
        "errors"
        "sync"
)

var (
      DB = newDB() // DB is reference of database
 
      ErruserAlreadyExists = errors.New("user name already exists") 
      ErruserDoesNotExists = errors.New("User does not exists")

//      m = map[string]string

)

//IN memory database
type Store struct {
     rwm *sync.RWMutex  // protected by mutex so that two go routines cant write to underlying maps at same time
     m   map[string]string
}

func newDB() *Store {
     return &Store {
            rwm:  &sync.RWMutex{},
              m:  make(map[string]string),
            }
}


func NewUser(un string, pw string) error {
     err := exists(un)
     if err != nil {
        return err
     }

     DB.rwm.Lock()
     defer DB.rwm.Unlock()

     bpass, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
     if err != nil {
        return err
     }

     DB.m[un] = string(bpass)

     fmt.Printf("Map value is %v\n", DB.m) 

     return nil 

}

func exists(un string) error {
     DB.rwm.Lock()
     defer DB.rwm.Unlock()

     if DB.m[un] != "" {
        return ErruserAlreadyExists
     }
    
     return nil
}
       

func AuthenticateUser(un string, pw string) error {
    
     DB.rwm.Lock()
     defer DB.rwm.Unlock()

     hashPwd, ok := DB.m[un]
     if !ok {
         return ErruserDoesNotExists
     }

    err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pw))
    if err == nil {
       fmt.Println("Authentication done sucessfully")
    }
    return err
}        
  

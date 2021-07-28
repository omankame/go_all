package resetemail
       
import (
         "fmt"
         "golang.org/x/crypto/bcrypt"
//         "sync"
         "errors"
)

var (
      DB = newDB()
      ErrUserAlreadyExists = errors.New("User alredy Exists")
      ErrUserDoesNotExists = errors.New("Given user name is not present in database")
)

type Store struct {
     m  map[string]string
}

func newDB() *Store {
     return &Store {
            m: make(map[string]string),
           }
}

func CreateUser(un string, pw string) error {
     err := Exists(un)
     if err != nil {
        return err
     }

    hashPwd, err :=  bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
    if err != nil {
       return err
    }

   DB.m[un] = string(hashPwd)

    return nil
}

func Exists(un string) error {
    
     if DB.m[un] != "" {
        return ErrUserAlreadyExists
     }

     return nil

}


func AuthenticateUser(un string, pw string) error {
     hashPwd, ok := DB.m[un]
     if !ok {
        return ErrUserDoesNotExists
     }


     err :=  bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pw)) 
     if err != nil {
        return err
     }

     fmt.Println("Authentication doen sucessfully")
     return nil
}

func OverrideOldPassword(un string, pw string ) error {

      hashPwd, err :=  bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
      if err != nil {
         return err
      }

      DB.m[un] = string(hashPwd)

    return nil
}

     

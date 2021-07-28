package erhandle

import (
        "net/http"
        "log"
        "os"
        "time"
        "errors"
)

var (
      ErrInvalidID  = errors.New("Invalid ID")
      ErrInvalidEmailID = errors.New("InvalidEmailID")
)

func CreateLog(lFile string) *log.Logger {
     file, err := os.OpenFile(lFile+".log", os.O_CREATE | os.O_WRONLY | os.O_APPEND , 0644)
     if err != nil {
        panic(err)
     }      

    logr :=  log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile) 
    return logr
}
 
func Recover(panicerr http.HandlerFunc) http.Handler {
     return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
            defer func() { 
                  err := recover() 
                  if err != nil {
                     switch err {
                        case ErrInvalidEmailID:
                             http.Error(w, ErrInvalidEmailID.Error(), http.StatusUnauthorized)
                        case ErrInvalidID:
                             http.Error(w, ErrInvalidID.Error(), http.StatusUnauthorized)
                        default:
                             http.Error(w, "Unknown Error, recovered from panic", http.StatusInternalServerError)
                     }
                  }
               }()
              panicerr.ServeHTTP(w, r)
           })
}
     

func Time(logr *log.Logger, next http.HandlerFunc) http.Handler {
     return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
                 start := time.Now()
                 next.ServeHTTP(w, r)
                 elapsed := time.Since(start)
                 logr.Println(elapsed)                
           })


}
 

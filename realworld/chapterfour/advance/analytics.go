package advance

import (
         "net/http"
         "log"
         "errors"
         "time"
         "os"
         "golang.org/x/net/context" 
)

var (
     ErrInvalidID      = errors.New("InvalidID")
     ErrInvalidEmailID = errors.New("InvalidEmailID")
) 

func CreateLog(str string) *log.Logger {
     file, err := os.OpenFile(str+".log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
     if err != nil {
        panic(err)
     }

     logger := log.New(file, "", log.Ldate| log.Ltime| log.Lshortfile) 
     return logger

}

func Next(logger *log.Logger, next http.HandlerFunc) http.Handler {
     return http.HandlerFunc( func (w http.ResponseWriter, req *http.Request) {
            start := time.Now()
            next.ServeHTTP(w, req)
            elapsed := time.Since(start)
            logger.Println(elapsed)
     })

}

func Recover(next http.HandlerFunc) http.Handler {
         return http.HandlerFunc( func (w http.ResponseWriter, req *http.Request) {
         defer func() {
               err := recover()
                if err != nil {
                  switch err {
                   case ErrInvalidID:
                      http.Error( w, ErrInvalidID.Error(), http.StatusUnauthorized) 
                   case ErrInvalidEmailID:
                      http.Error( w, ErrInvalidEmailID.Error(), http.StatusUnauthorized)
                   default:
                      http.Error(w , "Unknow error", http.StatusInternalServerError)
                 }
               }
         }()
      next.ServeHTTP(w, req)
      })
}

type PassContext func(ctx context.Context, w http.ResponseWriter, req *http.Request)

func (fn PassContext) ServeHTTP(w http.ResponseWriter, req *http.Request) {
      ctx := context.WithValue(context.Background(), "foo", "bar")
      fn(ctx, w, req)
} 

package logger

import (
         "net/http"
         "os"
         "time"
         "log"
)

func CreateLog(fname string) *log.Logger {
     file, err := os.OpenFile(fname+".log", os.O_CREATE| os.O_WRONLY| os.O_APPEND, 0666)
     if err != nil {
        panic(err)
     }

     logr := log.New(file, "", log.Ldate| log.Ltime | log.Lshortfile) 
     return logr

}


func TIME(logr *log.Logger, next http.HandlerFunc) http.Handler {
       return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                            start := time.Now()
                            next.ServeHTTP(w,r)
                            elapsed := time.Since(start)
                            logr.Println(elapsed)

                          })
}


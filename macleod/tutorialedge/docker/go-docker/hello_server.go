// https://www.callicoder.com/docker-golang-image-container-example/

package main

import (
        "fmt"
        "log"
        "gopkg.in/natefinch/lumberjack.v2"
        "github.com/gorilla/mux"
        "net/http"
        "os"
        "os/signal"
        "syscall"
        "context"
        "time"
)

func main() {

//    log_loct := os.Getenv("Log_File")
//    LOG_FILE_LOCATION := os.Getenv("LOGNAME")
//    fmt.Println(LOG_FILE_LOCATION) 

        r := mux.NewRouter()

        r.HandleFunc("/", handler)

        srv := &http.Server {
                 Handler:  r,
                 Addr:    ":8080",
        }

           LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	     if LOG_FILE_LOCATION != "" {        
             log.SetOutput(&lumberjack.Logger{
                        Filename:   LOG_FILE_LOCATION,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
          }
         
        go  func() {
            log.Println("Starting Server")
            err := srv.ListenAndServe() 
            if err != nil {
               log.Fatal(err)
            }
          }()

          waitForShutdown(srv)

}


func handler(w http.ResponseWriter, req *http.Request) {
       query := req.URL.Query()
       str   := query.Get("name")
       if str == " " {
          str = "Guest"
       }

      log.Printf("Received request for %s\n", str)
      w.Write([]byte(fmt.Sprintf("Hello %s\n",str)))
} 
              
func waitForShutdown(srv *http.Server) {
     interruptChan := make(chan os.Signal, 1)
     signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
      
     <-interruptChan

     ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
     defer cancel()
     srv.Shutdown(ctx)
     
    log.Println("Shutting Down")
    os.Exit(1)
}

    
        
            


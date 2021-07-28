package main

import (
        "fmt"
        "net/http"
        "github.com/gorilla/websocket"
        "log"
)

var upgrader = websocket.Upgrader {
               ReadBufferSize: 1024,   // size in bytes for every clinet connection
               WriteBufferSize: 1024,
} 


func main() {
     fmt.Println("Hello World")

     setupRoutes()
     http.ListenAndServe(":8080", nil)
}

func setupRoutes() {

     http.HandleFunc("/", homePage)
     http.HandleFunc("/ws", wsEndPoint)
}

func homePage(w http.ResponseWriter, req *http.Request) {
     fmt.Fprint(w, "Home Page") 

}

func wsEndPoint(w http.ResponseWriter, req *http.Request) {
//     fmt.Fprint(w, "Hello World, First Websocket Page")

       upgrader.CheckOrigin = func(req *http.Request) bool { 
                                                         //whether or not an incoming request from a different domain is allowed to connt
                              return true
                             }

       ws, err := upgrader.Upgrade(w, req, nil)  //ws is connection pinter websocket
       if err != nil {
          log.Println(err)
       }

       // helpful log statement to show connections
       log.Println("Client Connected")
       _  = ws.WriteMessage(1, []byte("Hi Client!"))
       reader(ws)  // define a reader which will listen for new messages being sent to our WebSocket endpoint
}

func reader(conn *websocket.Conn) {
     
     for {
          messageType, p, err :=  conn.ReadMessage()
          if err != nil {
             log.Println(err)
             return
          }

          fmt.Println(string(p))

         err =  conn.WriteMessage(messageType, p)
         if err != nil {
            log.Println(err)
            return
         }  
     }

}

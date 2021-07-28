package main

import (
         "log"
         "net/http"
         "github.com/gorilla/websocket"
)

var (
      clients = make(map[*websocket.Conn]bool)  //connected clients
      broadcast = make(chan Message)   //broadcast channel
      upgrader = websocket.Upgrader{}
)

type Message struct {
     Email string    `json:"email"` 
     Username string `json:"username"`
     Message  string `json:message"`
}

func main() {

     fs := http.FileServer(http.Dir("./public"))
     http.Handle("/", fs)
     http.HandleFunc("/ws", handleConnections)

     go handleMessages()

     log.Println("Starting WebServer on :8080")
     http.ListenAndServe(":8080", nil)

}

func handleConnections(w http.ResponseWriter, req *http.Request) {
     ws, err := websocket.Upgrader.Upgrade(w, r, nil)
     if err != nil {
        log.Fatal(err)
     }

     defer ws.Close()

     clients[ws] = true
     
     for {
          var msg Message 
          err := ws.ReadJson(&msg)
          if err != nil {
             log.Printf("error: %v", err)
             delete(clients, ws)
          }

          broadcast <- msg
    }

}

func handleMessages() {
     for {
          msg := <- broadcast

          for client := range clients {
                 err := client.WriteJSON(msg)
                 if err != nil {
                     log.Printf("error: %v", err)
                     client.Close()
                     delete(clients, client)
                 }

           }

    }

} 
    

}      
      
    




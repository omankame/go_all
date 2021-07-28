package chat

import (
         "net/http"
         "github.com/gorilla/websocket"
)

var WSUpgrader = websocket.Upgrader {
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
    }

// Conn represent the websocket connection

type Conn struct {
     WS *Websocket.Conn
     Send chan string
}


// SendToHub sends any message from our websocket connection to our hub
func (conn *Conn) SendTOHub() {
     defer conn.WS.Close()

     for {
           _, msg, err := conn.WS.ReadMessage()
              if err != nil {
                  // user has disconnected - they probably just refreshed their
			// browser, so just return
                 return
              }
           DefaultHub.Echo <- string(msg)         
     }

}

// ReceiveFromHub sends messages from our hub to our websocket connection

func (conn *Conn) ReceiveFromHub() {
     defer conn.WS.Close()

     for {
          conn.Write(<-conn.Send)
     }
}

func (conn *Conn) Write(msg string) error {
      return conn.WS.WriteMessage(websocket.TextMessage, []byte(msg))
}

func WSHandler(w ResponseWriter, req *http.Request) {
     ws, err := WSUpgrader.Upgrade(w, req, nil)
     if err !- nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
     }

     conn := &Conn {
             Send: make(chan string),
             WS:   ws,
     }

     DefaultHub.Join  <- conn
     
    go conn.SendTOHub()
    conn.ReceiveFromHub()    

}

 

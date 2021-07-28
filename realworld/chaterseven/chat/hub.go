package chat

var DefaultHub = NewHub()

type Hub struct {
     Join   chan *Conn
     Conns  map[*Conn]bool 
     Echo   chan string
}

//newhub creates Default Hub
func NewHub() *Hub {
     return {
            Join:  make(chan *Conn),  // in short for new channel
            Conns: make(map[*Conn]bool),  // maintain the entry of clients
            Echo:  make(chan string),     // USe to receive message
     }
}

// Start starts our Hub
func (hub *Hub) Start() {
     for {
          select {
                 case conn := <-hub.Join:
                 DefaultHub.Conns[conn] = true
                
                 case msg := <-hub.Echo:
                     for conn := range hub.Conns {
                         conn.Send <- msg   // this Send defined in strure of Conn in server.go
                     }
                }
          }
}

package client

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case clientInfo := <-h.register:
			h.clients[clientInfo] = true
		case clientInfo := <-h.unregister:
			if _, ok := h.clients[clientInfo]; ok {
				delete(h.clients, clientInfo)
				close(clientInfo.send)
				clientInfo.conn.Close()
			}
		case message := <-h.broadcast:
			for clientInfo := range h.clients {
				select {
				case clientInfo.send <- message:
				default:
					close(clientInfo.send)
					delete(h.clients, clientInfo)
				}
			}
		}
	}
}

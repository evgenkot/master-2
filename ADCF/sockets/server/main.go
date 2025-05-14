package main

import (
	"log"
	"net/http"
	"sync"
	"time"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Message struct {
	Type      string    `json:"type"`
	From      string    `json:"from,omitempty"`
	Content   string    `json:"content,omitempty"`
	Room      string    `json:"room,omitempty"`
	Users     []string  `json:"users,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

type Client struct {
	conn     *websocket.Conn
	name     string
	room     string
	server   *Server
}

type Room struct {
	clients map[*Client]bool
	mutex   sync.Mutex
}

type Server struct {
	rooms map[string]*Room
	mutex sync.Mutex
}

func NewServer() *Server {
	return &Server{
		rooms: make(map[string]*Room),
	}
}

func (s *Server) addClientToRoom(client *Client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	room, exists := s.rooms[client.room]
	if !exists {
		room = &Room{
			clients: make(map[*Client]bool),
		}
		s.rooms[client.room] = room
	}

	room.mutex.Lock()
	room.clients[client] = true
	room.mutex.Unlock()
}

func (s *Server) removeClientFromRoom(client *Client) {
	s.mutex.Lock()
	room, exists := s.rooms[client.room]
	s.mutex.Unlock()

	if !exists {
		return
	}

	room.mutex.Lock()
	delete(room.clients, client)
	room.mutex.Unlock()

	if len(room.clients) == 0 {
		s.mutex.Lock()
		delete(s.rooms, client.room)
		s.mutex.Unlock()
	}
}

func (s *Server) broadcast(message Message, roomName string) {
	s.mutex.Lock()
	room, exists := s.rooms[roomName]
	s.mutex.Unlock()

	if !exists {
		return
	}

	room.mutex.Lock()
	defer room.mutex.Unlock()

	for client := range room.clients {
		err := client.conn.WriteJSON(message)
		if err != nil {
			log.Printf("Error sending message: %v", err)
			client.conn.Close()
			s.removeClientFromRoom(client)
		}
	}
}

func (s *Server) handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	var initMsg Message
	err = conn.ReadJSON(&initMsg)
	if err != nil || initMsg.Type != "join" {
		log.Println("Invalid join message")
		return
	}

	client := &Client{
		conn:   conn,
		name:   initMsg.From,
		room:   initMsg.Room,
		server: s,
	}

	client.server.addClientToRoom(client)
	defer client.server.removeClientFromRoom(client)

	joinMsg := Message{
		Type:      "system",
		Content:   initMsg.From + " has joined the room",
		Timestamp: time.Now(),
	}
	client.server.broadcast(joinMsg, client.room)

	client.sendUserList()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			break
		}

		switch msg.Type {
		case "message":
			msg.Timestamp = time.Now()
			client.server.broadcast(msg, client.room)
		case "list_users":
			client.sendUserList()
		}
	}

	leaveMsg := Message{
		Type:      "system",
		Content:   client.name + " has left the room",
		Timestamp: time.Now(),
	}
	client.server.broadcast(leaveMsg, client.room)
}

func (c *Client) sendUserList() {
	c.server.mutex.Lock()
	room, exists := c.server.rooms[c.room]
	c.server.mutex.Unlock()

	if !exists {
		return
	}

	users := make([]string, 0, len(room.clients))
	room.mutex.Lock()
	for client := range room.clients {
		users = append(users, client.name)
	}
	room.mutex.Unlock()

	msg := Message{
		Type:  "user_list",
		Users: users,
	}
	c.conn.WriteJSON(msg)
}

func main() {
	server := NewServer()
	http.HandleFunc("/ws", server.handleConnection)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	server *Server
}

func NewHandler(server *Server) *Handler {
	return &Handler{
		server: server,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) ServeWs(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	name, ok := c.Request.URL.Query()["name"]

	if !ok || len(name[0]) < 1 {
		log.Printf("Invalid name param: %v", err)
		return
	}

	client := newClient(conn, h.server, name[0])

	fmt.Println("New Client joined the server!")
	fmt.Println(client)

	go client.writePump()
	go client.readPump()

	h.server.register <- client

}

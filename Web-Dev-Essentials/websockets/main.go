package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// allow all origins for simplicity
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to initiate connection"})
		return
	}

	log.Println("Client connected")

	go hanldeConnection(conn)
}

func hanldeConnection(conn *websocket.Conn) {
	defer conn.Close()
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("error reading message: ", err.Error())
			break
		}
		log.Printf("message received: %v", message)

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("error writing message: ", err.Error())
			break
		}
	}
}

func main() {
	r := gin.Default()

	r.GET("/ws", handleWebSocket)
	r.Run(":3001")
}

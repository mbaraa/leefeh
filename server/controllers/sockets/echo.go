package sockets

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type EchoSocket struct {
	remotes []*websocket.Conn
}

func NewEchoSocket() *EchoSocket {
	return &EchoSocket{}
}

func (e *EchoSocket) Bind(app *fiber.App) {
	socketGroup := app.Group("/ws/echo")
	socketGroup.Get("/", websocket.New(e.handleEcho))
}

func (e *EchoSocket) handleEcho(conn *websocket.Conn) {
	defer conn.Close()

	e.remotes = append(e.remotes, conn)

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(string(data))
		for _, member := range e.remotes {
			// if member.RemoteAddr().String() != conn.RemoteAddr().String() {
			err = member.WriteMessage(websocket.TextMessage, data)
			// }
			if err != nil {
				break
			}
		}
	}

}

package sockets

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client interface {
	SendMessage([]byte)
	GetID() uuid.UUID
	GetNickname() string
	Close()
}

type client struct {
	ID       uuid.UUID
	Nickname string
	Socket   *websocket.Conn

    Receive chan []byte

    Lobby Lobby
}

var Upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func NewClient(conn *websocket.Conn, l Lobby, nickname string) Client {
	c := &client{
		ID:       uuid.New(),
		Nickname: nickname,
		Socket:   conn,
		Receive:  make(chan []byte, Upgrader.ReadBufferSize),
		Lobby:    l,
	}

    l.JoinLobby(c)

    defer func() {
        l.LeaveLobby(c)
    }()

    go c.Write()
    c.Read()

    return c
}

func (c *client) Read() {
    defer c.Socket.Close()
    for {
        _, msg, err := c.Socket.ReadMessage()
        fmt.Println("Received message: ", string(msg))
        if err != nil {
            _ = fmt.Errorf("there was a error when reading message for Client: %w", err)
            return
        }
        c.Lobby.ForwardMessage(Message{
            SocketID: c.ID,
            Data:     msg,
        })
    }
}

func (c *client) Write() {
    defer c.Socket.Close()
    for msg := range c.Receive {
        fmt.Println("Writing message: ", string(msg))
        err := c.Socket.WriteMessage(websocket.TextMessage, msg)
        if err != nil {
            return
        }
    }
}

func (c *client) SendMessage(msg []byte) {
    c.Receive <- msg
}

func (c *client) Close() {
    close(c.Receive)
}

func (c *client) GetID() uuid.UUID {
    return c.ID
}

func (c *client) GetNickname() string {
	return c.Nickname
}

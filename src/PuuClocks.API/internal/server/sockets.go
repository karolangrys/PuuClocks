package server

import (
	"puuclocks/internal/log"
	"puuclocks/internal/repository"
	"puuclocks/internal/server/sockets"
	"puuclocks/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/gorilla/websocket"
)

type socketServer struct {
	service      service.Service
	databases    repository.Databases
	lobbyManager sockets.LobbyManager
}

type SocketServerParameters struct {
	Service      service.Service
	Databases    repository.Databases
	LobbyManager sockets.LobbyManager
}

func NewSocketServer(p SocketServerParameters) *socketServer {
	return &socketServer{
		service:      p.Service,
		databases:    p.Databases,
		lobbyManager: p.LobbyManager,
	}
}

func (s socketServer) JoinLobby(c *gin.Context) {
	var conn *websocket.Conn
	var parsedID uuid.UUID
	var err error

	conn, err = sockets.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Log.Errorln(err)
		return
	}
	defer conn.Close()
	id := c.Param("id")
	nickname := c.Query("nickname")
	if nickname == "" {
		if err = conn.WriteJSON(map[string]string{
			"message": "User not passed nickname",
		}); err != nil {
			log.Log.Errorln(err)
		}
		return
	}

	parsedID, err = uuid.Parse(id)
	if err != nil {
		if err = conn.WriteJSON(map[string]string{
			"message": "User not passed lobby UUID",
		}); err != nil {
			log.Log.Errorln(err)
		}
	} else {
		l := s.lobbyManager.FindLobby(parsedID)
		if l == nil {
			if err = conn.WriteJSON(map[string]string{
				"message": "Lobby not found",
			}); err != nil {
				log.Log.Errorln(err)
			}
		} else {
			sockets.NewClient(conn, l, nickname)
			if err = conn.WriteJSON(map[string]string{
				"message": "User connected",
			}); err != nil {
				log.Log.Errorln(err)
			}
		}
	}
}

func (s socketServer) RegisterSocketHandlers(router gin.IRouter) {
	router.GET("/join-lobby/:id", s.JoinLobby)
}

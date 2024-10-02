package server

import (
	"net/http"

	"puuclocks/internal/repository"
	"puuclocks/internal/server/sockets"
	"puuclocks/internal/service"

	"github.com/gin-gonic/gin"
)

type restServer struct {
	service      service.Service
	databases    repository.Databases
	lobbyManager sockets.LobbyManager
}

type RestServerParameters struct {
	Service      service.Service
	Databases    repository.Databases
	LobbyManager sockets.LobbyManager
}

func NewRestServer(p RestServerParameters) *restServer {
	return &restServer{
		service:      p.Service,
		databases:    p.Databases,
		lobbyManager: p.LobbyManager,
	}
}

func (s *restServer) Ping(c *gin.Context) {
	if err := s.service.Health(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "DOWN"})
		return
	}

	if err := s.databases.Health(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "DOWN"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}

func (s *restServer) CreateLobby(c *gin.Context) {
	lobby := s.lobbyManager.CreateLobby(s.service)

	c.JSON(http.StatusOK, gin.H{
		"lobbyID": lobby.GetID(),
	})
}

package main

import (
	"net/http"
	"puuclocks/internal/infrastructure"
	"puuclocks/internal/log"
	"puuclocks/internal/repository"
	"puuclocks/internal/server"
	"puuclocks/internal/service"
	"puuclocks/internal/sockets"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	gen_openapi "puuclocks/gen"
)

func main() {
	log.InitLogger()

	dbCfg := repository.DatabasesConfig{
		RedisConfig: repository.RedisConfig{
			Addr: "redis:6379",
		},
		MySQLConfig: infrastructure.MySQLConfig{
			DBName: "mysql",
			Path:   "root:root@tcp(mysql:3306)/puuclocks",
		},
	}

	databases, err := repository.NewDatabases(&dbCfg)
	if err != nil {
		panic(err)
	}

	s := service.NewService()
	lobbyManager := sockets.NewLobbyManager()

	rest := server.NewRestServer(server.RestServerParameters{
		Service:      s,
		Databases:    databases,
		LobbyManager: lobbyManager,
	})
	socket := server.NewSocketServer(server.SocketServerParameters{
		Service:      s,
		Databases:    databases,
		LobbyManager: lobbyManager,
	})

	r := gin.Default()

	gen_openapi.RegisterHandlers(r, rest)
	socket.RegisterSocketHandlers(r)

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost"
		},
		MaxAge: 12 * time.Hour,
	}))

	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: time.Second,
	}

	log.Log.DPanicln(httpServer.ListenAndServe())
}

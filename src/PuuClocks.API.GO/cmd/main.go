package main

import (
	"net/http"
	"puuclocks/internal/infrastructure"
	"puuclocks/internal/log"
	"puuclocks/internal/repository"
	"puuclocks/internal/server"
	"puuclocks/internal/server/sockets"
	"puuclocks/internal/service"
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

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		AllowWebSockets:  true,
	}))

	r.Use(CORSMiddleware())

	gen_openapi.RegisterHandlers(r, rest)
	socket.RegisterSocketHandlers(r)

	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: time.Second,
	}

	log.Log.DPanicln(httpServer.ListenAndServe())
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:2137")

		c.Next()
	}
}

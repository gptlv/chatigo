package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gptlv/chatigo/server/internal/delivery/restapi"
	"github.com/gptlv/chatigo/server/internal/delivery/ws"
)

var r *gin.Engine

func InitRouter(uh *restapi.UserHandler, wsh *ws.Handler) {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/signup", uh.CreateUser)
	r.POST("/login", uh.Login)
	r.GET("/logout", uh.Logout)

	r.GET("/ws", wsh.ServeWs)

}

func Start(addr string) error {
	return r.Run(addr)
}

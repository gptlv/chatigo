package main

import (
	"log"

	"github.com/gptlv/chatigo/server/db"
	"github.com/gptlv/chatigo/server/internal/user"
	"github.com/gptlv/chatigo/server/internal/ws"
	"github.com/gptlv/chatigo/server/router"
)

func main() {

	dbConn, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("DB connection error: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	hubHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, hubHandler)
	router.Start("0.0.0.0:8080")

}

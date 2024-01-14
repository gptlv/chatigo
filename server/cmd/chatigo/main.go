package main

import (
	"log"

	"github.com/gptlv/chatigo/server/internal/delivery/restapi"
	"github.com/gptlv/chatigo/server/internal/delivery/router"
	"github.com/gptlv/chatigo/server/internal/delivery/ws"
	db "github.com/gptlv/chatigo/server/internal/repository/postgres"
	postgres "github.com/gptlv/chatigo/server/internal/repository/postgres/user"
	usecase "github.com/gptlv/chatigo/server/internal/usecase/user"
)

func main() {

	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("DB connection error: %s", err)
	}

	userRepo := postgres.NewRepository(dbConn)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := restapi.NewUserHandler(userUsecase)

	wsServer := ws.NewServer()
	wsHandler := ws.NewHandler(wsServer)
	go wsServer.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:1337")

}

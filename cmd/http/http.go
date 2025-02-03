package main

import (
	"log"

	"github.com/freitzzz/gameboy-db-api/internal/data"
	"github.com/freitzzz/gameboy-db-api/internal/database"
	"github.com/freitzzz/gameboy-db-api/internal/http"
	"github.com/freitzzz/gameboy-db-api/internal/logging"
	"github.com/freitzzz/gameboy-db-api/internal/service"
)

func main() {
	db, err := database.Open("/home/freitas/Workspace/Projects/freitzzz/gameboy-db-api/database/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	r := data.NewDbGamesRepository(db)
	s := service.NewGamesService(r)

	logging.AddLogger(logging.NewConsoleLogger())

	hs := http.New().
		WithHostPort("localhost", "8080").
		WithServiceContainer(http.ServiceContainer(s)).
		Build()

	defer hs.Close()
	if err := hs.Start(); err != nil {
		logging.Fatal("failed to start http server, %v", err)
	}
}

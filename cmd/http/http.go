package main

import (
	"github.com/freitzzz/gameboy-db-api/internal/build"
	"github.com/freitzzz/gameboy-db-api/internal/data"
	"github.com/freitzzz/gameboy-db-api/internal/database"
	"github.com/freitzzz/gameboy-db-api/internal/env"
	"github.com/freitzzz/gameboy-db-api/internal/http"
	"github.com/freitzzz/gameboy-db-api/internal/logging"
	"github.com/freitzzz/gameboy-db-api/internal/service"
)

func init() {
	if build.Release() && !build.Verbose() {
		logging.DisableDebugLogs()
	}

	logging.AddLogger(logging.NewConsoleLogger())
}

func main() {
	env := env.Env

	if env.DBPath == nil {
		logging.Fatal("cannot start server since db path is unknown")
	}

	db, err := database.Open(*env.DBPath)
	if err != nil {
		logging.Fatal("failed to connect to database, (%v)", err)
	}

	r := data.NewDbGamesRepository(db)
	s := service.NewGamesService(r)

	hs := http.Builder().
		WithHostPort(env.ServerHost, env.ServerPort).
		WithSelfSignedCertificate(env.ServerTLSCert, env.ServerTLSKey).
		WithVirtualPath(env.ServerVirtualPath).
		WithServiceContainer(http.ServiceContainer(s)).
		Build()

	defer hs.Close()
	defer db.Close()
	if err := hs.Start(); err != nil {
		logging.Fatal("failed to start http server, %v", err)
	}
}

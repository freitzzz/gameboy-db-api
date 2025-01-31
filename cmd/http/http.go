package main

import (
	"fmt"
	"log"

	"github.com/freitzzz/gameboy-db-api/internal/database"
)

func main() {
	db, err := database.Open("/home/freitas/Workspace/Projects/freitzzz/gameboy-db-api/database/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	tbls, err := db.Tables()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tbls: %v\n", tbls)
}

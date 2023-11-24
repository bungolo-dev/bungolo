package main

import (
	bungolow "github.com/bungolow-dev/bungolow/pkg/application"
	"github.com/bungolow-dev/bungolow/pkg/infrustucture"
)

func main() {

	db := infrustucture.CreateDatabase()
	err := db.Initialize()
	if err != nil {
		bungolow.Logger.Error("Failed to initialize the database", err)
	}

	room, err := db.CreateRoom("Eric's Room", "http://127.0.0.1:8081/images/room.png")
	if err != nil {
		bungolow.Logger.Error("Failed to create new room", err)

	}

	bungolow.Logger.Information("Created new room %s", room)

}

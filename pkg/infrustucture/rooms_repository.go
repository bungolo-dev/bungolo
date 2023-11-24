package infrustucture

import (
	bungolow "github.com/bungolow-dev/bungolow/pkg/application"
	"github.com/bungolow-dev/bungolow/pkg/domain/room"
)

func (database *Database) CreateRoom(name string, image string) (room.RoomId, error) {

	id := room.CreateNewId()
	db, err := database.Open()
	if err != nil {
		bungolow.Logger.Error("Failed to open database %s", err)
		return "", err
	}
	defer db.Close()

	insertDataSQL := `
		INSERT INTO rooms (id, name, description, image)
		VALUES (?, ?, ?, ?);
	`
	_, e := db.Exec(insertDataSQL, id, name, "description", image)
	if e != nil {
		bungolow.Logger.Error("Failed to insert room into database %s", err)
		return "", e
	}

	bungolow.Logger.Information("Created new Room in Database %s", id)
	return id, nil
}

func (database *Database) EditRoom(id room.RoomId, name string, image string) error {
	db, err := database.Open()
	if err != nil {
		bungolow.Logger.Error("Failed to open database %s", err)
		return err
	}
	defer db.Close()

	updateDataSQL := `
		UPDATE INTO rooms (id, name, description, image)
		VALUES (?, ?, ?, ?);
	`
	_, e := db.Exec(updateDataSQL, id, name, "description", image)
	if e != nil {
		bungolow.Logger.Error("Failed to update room in database %s", err, id)
		return e
	}
	return nil
}

func (database *Database) DeleteRoom(id room.RoomId) error {
	db, err := database.Open()
	if err != nil {
		bungolow.Logger.Error("Failed to open database %s", err)
		return err
	}
	defer db.Close()

	deleteDataSQL := `
		DELETE FROM rooms WHERE id = ?;
	`
	_, e := db.Exec(deleteDataSQL, id)
	if e != nil {
		bungolow.Logger.Error("Failed to delete room from database %s", err, id)
		return e
	}
	return nil
}

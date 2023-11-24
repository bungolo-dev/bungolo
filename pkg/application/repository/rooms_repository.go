package repository

import "github.com/bungolow-dev/bungolow/pkg/domain/room"

type RoomsRepository interface {
	CreateRoom(name string, image string) (room.RoomId, error)
	EditRoom(id room.RoomId, name string, image string) error
	DeleteRoom(id room.RoomId) error
}

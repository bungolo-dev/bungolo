package repository

import "github.com/bungolow-dev/bungolow/pkg/application/rooms"

type DbConnection interface {
	Initialize() error
	rooms.RoomsRepository
}

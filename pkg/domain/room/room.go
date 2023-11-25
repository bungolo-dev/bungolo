package room

import (
	"net/url"

	"github.com/google/uuid"

	"github.com/bungolow-dev/bungolow/pkg/domain/destinations"
	"github.com/bungolow-dev/bungolow/pkg/domain/sources"
)

type Room struct {
	Id          RoomId
	Name        string
	Description string
	Image       url.URL

	Sources  *[]sources.Source
	Displays *[]destinations.Destination
}

type RoomId string

func CreateNewId() RoomId {
	return RoomId(uuid.New().String())
}

func CreateExistingId(id string) RoomId {
	//VALIDATION HERE
	return RoomId(id)
}

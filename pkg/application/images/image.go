package images

import "github.com/google/uuid"

type Image struct {
	Id   ImageId
	Name string
	Path string
}

type ImageId string

func CreateNewId() ImageId {
	return ImageId(uuid.New().String())
}

func CreateExistingId(id string) ImageId {
	return ImageId(id)
}

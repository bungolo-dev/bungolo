// Serves images to user interfaces
//
// DB Tables
// | Image ID | Image Name | Relative Path |
//
// The images repository stores physical graphics fles mapped to image IDs.
// All image files paths are stored in directory defined in the bungolow enviroment
// confguration.
package images

import (
	"fmt"
	"log"
	"os"

	bungolow "github.com/bungolow-dev/bungolow/pkg/application"
)

var images []Image = []Image{}

func GetAllImages() []Image {

	if len(images) > 0 {
		return images
	}

	files, err := os.ReadDir(bungolow.GetEnvironment().ImgDir())
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		name := file.Name()
		images = append(images, Image{
			Id:   CreateNewId(),
			Name: name,
			Path: bungolow.GetEnvironment().GetImagePath(name),
		})
	}

	return images
}

func GetImageById(id ImageId) (Image, error) {
	for _, img := range images {
		if img.Id == id {
			return img, nil
		}
	}

	return Image{}, fmt.Errorf("FAILED TO FIND IMAGE %s", id)
}
func CreateImage(filename string, filebytes []byte) (Image, error) {
	img := Image{
		Name: filename,
		Path: bungolow.GetEnvironment().GetImagePath(filename),
		Id:   CreateNewId(),
	}

	fileError := os.WriteFile(bungolow.GetEnvironment().GetImagePath(filename), filebytes, os.ModeAppend)
	if fileError != nil {
		return Image{}, nil
	}

	images = append(images, img)
	return img, nil
}

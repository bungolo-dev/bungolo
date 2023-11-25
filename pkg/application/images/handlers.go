package images

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	bungolow "github.com/bungolow-dev/bungolow/pkg/application"
	"github.com/go-chi/chi/v5"
)

func UseImageHandlers(mux *chi.Mux) {

	mux.Get("/api/images", GetImages)
	mux.Get("/api/images/*", GetImage)
	mux.Post("/api/images", PostImage)
}

func GetImages(w http.ResponseWriter, r *http.Request) {
	imgs := GetAllImages()

	json, err := json.MarshalIndent(imgs, "", "	")

	if err != nil {
		bungolow.HandleError(w, 500, err)
	}

	w.Write(json)
}

func GetImage(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.String(), "/")
	data, err := GetImageById(ImageId(parts[len(parts)-1]))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	}

	img, err := os.Open(data.Path)
	if err != nil {
		bungolow.Logger.Error("Failed to load image %s", err)
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, img)
}

func PostImage(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32 << 20)

	file, header, err := r.FormFile("file")
	if err != nil {
		bungolow.HandleError(w, http.StatusBadRequest, err)
		return
	}
	defer file.Close()
	var buf bytes.Buffer
	_, cpError := io.Copy(&buf, file)
	if cpError != nil {
		bungolow.HandleError(w, http.StatusBadRequest, cpError)
		return
	}

	image, saveError := CreateImage(header.Filename, buf.Bytes())

	if saveError != nil {
		bungolow.HandleError(w, http.StatusBadRequest, saveError)
		return
	}

	buf.Reset()

	json, jsonError := json.MarshalIndent(image, "", "   ")
	if jsonError != nil {
		bungolow.HandleError(w, http.StatusInternalServerError, jsonError)
	}

	w.Write(json)

}

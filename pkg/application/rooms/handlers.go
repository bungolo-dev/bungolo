package rooms

import (
	"net/http"
	"strings"

	bungolow "github.com/bungolow-dev/bungolow/pkg/application"
	"github.com/go-chi/chi/v5"
)

func UseRoomsHandlers(mux *chi.Mux) {

	mux.Get("/api/rooms", GetRooms)
	mux.Get("/api/rooms/*", GetRoom)
	mux.Post("/api/rooms", PostRoom)
}

func GetRooms(w http.ResponseWriter, r *http.Request) {

	bungolow.Logger.Information("full path %+v", r.URL)
	user := r.Context().Value("user").(string)
	w.Write([]byte(user))
}

func GetRoom(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.String(), "/")
	bungolow.Logger.Information("full path %+v", parts[len(parts)-1])
	user := r.Context().Value("user").(string)
	w.Write([]byte(user))
}

func PostRoom(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(string)
	//var bytes []byte
	//r.Body.Read(bytes)
	w.Write([]byte(user))
}

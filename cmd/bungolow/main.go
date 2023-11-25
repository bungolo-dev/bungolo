package main

import (
	"net/http"

	bungolow "github.com/bungolow-dev/bungolow/pkg/application"
	"github.com/bungolow-dev/bungolow/pkg/application/images"
	"github.com/bungolow-dev/bungolow/pkg/application/rooms"
	"github.com/bungolow-dev/bungolow/pkg/infrustucture"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	bungolow.NewEnvironment("C://Temp/Bungolow", "C://Temp/Images")

	db := infrustucture.CreateDatabase()
	err := db.Initialize()
	if err != nil {
		bungolow.Logger.Error("Failed to initialize the database", err)
	}

	room, err := db.CreateRoom("Bungolow", "http://127.0.0.1:8081/images/room.png")
	if err != nil {
		bungolow.Logger.Error("Failed to create new room", err)

	}

	bungolow.Logger.Information("Created new room %s", room)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	r.Use(MyMiddleware)

	rooms.UseRoomsHandlers(r)
	images.UseImageHandlers(r)

	http.ListenAndServe(":3000", r)
}

func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

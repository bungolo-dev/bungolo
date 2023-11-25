package plugins

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func UsePluginHandlers(mux *chi.Mux) {

	mux.Post("/api/plugins", PostCreatePlugin)
	mux.Get("/api/plugins", QueryPlugins)
	mux.Get("/api/plugins/*", QueryPlugin)
	mux.Post("/api/plugins/init", PostInitialize)
	mux.Delete("/api/plugins/*", DeletePlugin)
}

func PostCreatePlugin(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var settings PluginConfig

	err = json.Unmarshal(body, &settings)
	if err != nil {
		http.Error(w, "Error decoding JSON body", http.StatusBadRequest)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pluginErr := LoadPlugin(settings)
	if pluginErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func QueryPlugins(w http.ResponseWriter, r *http.Request) {

	var data string
	for _, p := range devices {
		q := p.device.Query()
		data = data + "\n" + q
	}
	w.Write([]byte(data))
}

func QueryPlugin(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.String(), "/")
	id := parts[len(parts)-1]

	device := devices[id]
	if device == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(id))
		return
	}

	q := device.device.Query()
	w.Write([]byte(q))
}

func PostInitialize(w http.ResponseWriter, r *http.Request) {

	for _, p := range devices {
		settigns := map[string]string{"ip": "102.22.55.5"}
		p.device.Initialize(settigns)
	}
}

func DeletePlugin(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.String(), "/")
	id := parts[len(parts)-1]

	err := KillPlugin(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
}

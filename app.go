package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
}

//Initialize this is great func
func (a *App) Initialize(user, password, dbname string) {

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/volumes", a.getVolumes).Methods("GET")
}

//Run is
func (a *App) Run(addr string) {
	println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) getVolumes(w http.ResponseWriter, r *http.Request) {
	return
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

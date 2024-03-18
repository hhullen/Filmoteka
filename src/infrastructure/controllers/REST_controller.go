package controllers

import (
	as "application_service"
	"db"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger = log.New(os.Stdout, "REST CONTROLLER: ", log.LstdFlags)

type ControllerREST struct {
	mux     *http.ServeMux
	service as.Filmoteka
}

func NewControllerREST(prefix string, mux *http.ServeMux) *ControllerREST {
	ctrl := &ControllerREST{
		mux:     mux,
		service: as.NewFilmoteka(db.NewMoMockDB()),
	}
	logger.Println("Service created")

	request_path := prefix + "/actors"
	mux.HandleFunc(request_path, ctrl.handleActors)
	logger.Println("Endpoint", request_path, "registered")
	return ctrl
}

func (me *ControllerREST) handleActors(w http.ResponseWriter, r *http.Request) {
	actor := &ActorDTO{}
	err := json.NewDecoder(r.Body).Decode(actor)
	if err != nil {
		logger.Println(r.Host, r.Method, r.URL, http.StatusBadRequest)
		logger.Fatal(err.Error())
		http.Error(w, "wrong body data", http.StatusBadRequest)
	}
	if r.Method == "GET" {
	} else if r.Method == "POST" {
	} else if r.Method == "PUT" {
	} else if r.Method == "DELETE" {
	}
	logger.Println(r.Host, r.Method, r.URL, http.StatusOK)
}

package controllers

import (
	as "application_service"
	db "postgres_db"
)

type ControllerREST struct {
	service as.Filmoteka
}

func NewControllerREST(prefix string) *ControllerREST {
	return &ControllerREST{service: as.NewFilmoteka(db.NewPostgredDB())}
}

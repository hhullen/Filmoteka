package controllers

import (
	dm "domain_model"
)

type ActorDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

func (me *ActorDTO) MapToDomain() (dm.Actors, error) {
	return dm.Actors{}, nil
}

func (me *ActorDTO) MapFromDomain(dm.Actors) {}

type FilmDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ReseaseDate string `json:"resease_date"`
	Rating      string `json:"rating"`
}

func (me *FilmDTO) MapToDomain() (dm.Films, error) {
	return dm.Films{}, nil
}

func (me *FilmDTO) MapFromDomain(dm.Films) {}

type ActorsFilmsRelationDTO struct {
	ActorID string
	FilmID  string
}

func (me *ActorsFilmsRelationDTO) MapToDomain() (dm.ActorsFilmsRelations, error) {
	return dm.ActorsFilmsRelations{}, nil
}

func (me *ActorsFilmsRelationDTO) MapFromDomain(dm.ActorsFilmsRelations) {}

type UserDTO struct {
	ID       string
	Role     string
	Login    string
	Password string
	JWT      string
	Expiring string
}

func (me *UserDTO) MapToDomain() (dm.Users, error) {
	return dm.Users{}, nil
}

func (me *UserDTO) MapFromDomain(dm.Users) {}

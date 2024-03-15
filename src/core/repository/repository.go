package repository

import (
	dm "domain_model"
)

type IDTO[Domain any] interface {
	MapToDomain() (Domain, error)
	MapFromDomain(Domain)
}

type IRepository interface {
	AddActor(actor dm.Actors) error
	UpdateActor(name string, actor dm.Actors) error
	DeleteActor(name string) error
	GetActor(name string) (dm.Actors, error)
	GetActorStarredFilms(name string) ([]dm.Films, error)
	AddFilm(film dm.Films) error
	AddFilmActors(film string, actors []dm.Actors) error
	DeleteFilmActors(film string, name string) error
	UpdateFilm(name string, film dm.Films) error
	DeleteFilm(name string) error
	GetAllFilms(order dm.SortOrder, column dm.SortColumn) ([]dm.Films, error)
	GetFilmsByNameSegment(segment string) ([]dm.Films, error)
	GetFilmsByActorNameSegment(segment string) ([]dm.Films, error)
	GetFilmCast(name string) ([]dm.Actors, error)
}

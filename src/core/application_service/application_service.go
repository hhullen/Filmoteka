package application_service

import (
	rep "repository"
)

type Filmoteka struct {
	repo rep.IRepository
}

func NewFilmoteka(repository rep.IRepository) Filmoteka {
	return Filmoteka{repo: repository}
}

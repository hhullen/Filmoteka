package application_service

import (
	dm "domain_model"
	"errors"
	rep "repository"
)

type Filmoteka struct {
	repo rep.IRepository
}

func NewFilmoteka(repository rep.IRepository) Filmoteka {
	return Filmoteka{repo: repository}
}

func (me *Filmoteka) AddActor(actor rep.IDTO[dm.Actors]) error {
	domain_actor, err := actor.MapToDomain()
	if err != nil {
		return errors.New(err.Error())
	}
	me.repo.AddActor(domain_actor)
	return nil
}

func (me *Filmoteka) UpdateActor(name string, actor rep.IDTO[dm.Actors]) error {
	domain_actor, err := actor.MapToDomain()
	if err != nil {
		return errors.New(err.Error())
	}
	return me.repo.UpdateActor(name, domain_actor)
}

func (me *Filmoteka) DeleteActor(name string) error {
	return me.repo.DeleteActor(name)
}

func (me *Filmoteka) GetActor(name string) (dm.Actors, error) {
	actor, err := me.repo.GetActor(name)
	if err != nil {
		return dm.Actors{}, errors.New(err.Error())
	}
	return actor, nil
}

func (me *Filmoteka) GetActorStarredFilms(name string) ([]dm.Films, error) {
	return me.repo.GetActorStarredFilms(name)
}

func (me *Filmoteka) AddFilm(film rep.IDTO[dm.Films]) error {
	domain_film, err := film.MapToDomain()
	if err != nil {
		return errors.New(err.Error())
	}
	return me.repo.AddFilm(domain_film)
}

func (me *Filmoteka) AddFilmActors(film string, actors []rep.IDTO[dm.Actors]) error {
	domain_actors := []dm.Actors{}
	for _, v := range actors {
		domain_actor, err := v.MapToDomain()
		if err != nil {
			return errors.New(err.Error())
		}
		domain_actors = append(domain_actors, domain_actor)
	}
	return me.repo.AddFilmActors(film, domain_actors)
}

func (me *Filmoteka) DeleteFilmActor(film string, name string) error {
	return me.repo.DeleteFilmActor(film, name)
}

func (me *Filmoteka) UpdateFilm(name string, film rep.IDTO[dm.Films]) error {
	domain_film, err := film.MapToDomain()
	if err != nil {
		return errors.New(err.Error())
	}
	return me.repo.UpdateFilm(name, domain_film)
}

func (me *Filmoteka) GetAllFilms(order dm.SortOrder, column dm.SortColumn) ([]dm.Films, error) {
	films, err := me.repo.GetAllFilms(order, column)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return films, nil
}

func (me *Filmoteka) GetFilmsByNameSegment(segment string) ([]dm.Films, error) {
	return me.repo.GetFilmsByNameSegment(segment)
}

func (me *Filmoteka) GetFilmsByActorNameSegment(segment string) ([]dm.Films, error) {
	return me.repo.GetFilmsByActorNameSegment(segment)
}

func (me *Filmoteka) GetFilmCast(name string) ([]dm.Actors, error) {
	return me.repo.GetFilmCast(name)
}

func (me *Filmoteka) DeleteFilm(name string) error {
	return me.repo.DeleteFilm(name)
}

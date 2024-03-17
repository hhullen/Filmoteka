package application_service

import (
	dm "domain_model"
	"errors"
	"log"
	"os"
	rep "repository"
)

type Filmoteka struct {
	logger *log.Logger
	repo   rep.IRepository
}

func NewFilmoteka(repository rep.IRepository) Filmoteka {
	logger := log.New(os.Stdout, "FILMOTEKA: ", log.LstdFlags)
	logger.Println("Service created")
	return Filmoteka{
		logger: logger,
		repo:   repository,
	}
}

func (me *Filmoteka) AddActor(actor rep.IDTO[dm.Actors]) error {
	domain_actor, err := actor.MapToDomain()
	if err != nil {
		me.logger.Println(err.Error())
		return errors.New(err.Error())
	}
	err = me.repo.AddActor(domain_actor)
	if err != nil {
		me.logger.Println(err.Error())
		return errors.New(err.Error())
	}
	me.logger.Println("Actor added", domain_actor)
	return nil
}

func (me *Filmoteka) UpdateActor(name string, actor rep.IDTO[dm.Actors]) error {
	domain_actor, err := actor.MapToDomain()
	if err != nil {
		me.logger.Println(err.Error())
		return errors.New(err.Error())
	}
	err = me.repo.UpdateActor(name, domain_actor)
	if err != nil {
		me.logger.Println(err.Error())
		return errors.New(err.Error())
	}
	me.logger.Println("Actor updated:", domain_actor)
	return nil
}

func (me *Filmoteka) DeleteActor(name string) error {
	err := me.repo.DeleteActor(name)
	if err != nil {
		me.logger.Println(err.Error())
		return errors.New(err.Error())
	}
	me.logger.Println("Actor deleted: \"" + name + "\"")
	return nil
}

func (me *Filmoteka) GetActor(name string) (dm.Actors, error) {
	actor, err := me.repo.GetActor(name)
	if err != nil {
		me.logger.Println(err.Error())
		return dm.Actors{}, errors.New(err.Error())
	}
	me.logger.Println("Actor got:", actor)
	return actor, nil
}

func (me *Filmoteka) GetActorStarredFilms(name string) ([]dm.Films, error) {
	films, err := me.repo.GetActorStarredFilms(name)
	if err != nil {
		me.logger.Println(err.Error())
		return nil, errors.New(err.Error())
	}
	me.logger.Println("Films got:", films)
	return films, nil
}

func (me *Filmoteka) AddFilm(film rep.IDTO[dm.Films]) error {
	domain_film, err := film.MapToDomain()
	if err != nil {
		me.logger.Println(err.Error())
		return errors.New(err.Error())
	}
	err = me.repo.AddFilm(domain_film)
	if err != nil {
		me.logger.Println(err.Error())
		return errors.New(err.Error())
	}
	me.logger.Println("Film added:", domain_film)
	return nil
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
	err := me.repo.AddFilmActors(film, domain_actors)
	if err != nil {
		me.logger.Println(err.Error())
		return errors.New(err.Error())
	}
	me.logger.Println("Film actor added:", domain_actors)
	return nil
}

func (me *Filmoteka) DeleteFilmActor(film string, name string) error {
	err := me.repo.DeleteFilmActor(film, name)
	if err != nil {
		me.logger.Println(err.Error())
		return errors.New(err.Error())
	}
	me.logger.Println("Film \"" + film + "\" actor \"" + name + "\" deleted")
	return nil
}

func (me *Filmoteka) UpdateFilm(name string, film rep.IDTO[dm.Films]) error {
	domain_film, err := film.MapToDomain()
	if err != nil {
		return errors.New(err.Error())
	}
	err = me.repo.UpdateFilm(name, domain_film)
	if err != nil {
		me.logger.Println(err.Error())
		return errors.New(err.Error())
	}
	me.logger.Println("Film updated:", domain_film)
	return nil
}

func (me *Filmoteka) GetAllFilms(order dm.SortOrder, column dm.SortColumn) ([]dm.Films, error) {
	films, err := me.repo.GetAllFilms(order, column)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	me.logger.Println("Films got:", films)
	return films, nil
}

func (me *Filmoteka) GetFilmsByNameSegment(segment string) ([]dm.Films, error) {
	films, err := me.repo.GetFilmsByNameSegment(segment)
	if err != nil {
		me.logger.Println(err.Error())
		return nil, errors.New(err.Error())
	}
	me.logger.Println("Films got:", films)
	return films, nil
}

func (me *Filmoteka) GetFilmsByActorNameSegment(segment string) ([]dm.Films, error) {
	films, err := me.repo.GetFilmsByActorNameSegment(segment)
	if err != nil {
		me.logger.Println(err.Error())
		return nil, errors.New(err.Error())
	}
	me.logger.Println("Films got:", films)
	return films, nil
}

func (me *Filmoteka) GetFilmCast(name string) ([]dm.Actors, error) {
	actors, err := me.repo.GetFilmCast(name)
	if err != nil {
		me.logger.Println(err.Error())
		return nil, errors.New(err.Error())
	}
	me.logger.Println("Film cast got:", actors)
	return actors, nil
}

func (me *Filmoteka) DeleteFilm(name string) error {
	err := me.repo.DeleteFilm(name)
	if err != nil {
		me.logger.Println(err.Error())
		return errors.New(err.Error())
	}
	me.logger.Println("Film deleted:", name)
	return nil
}

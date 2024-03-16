package db

import (
	dm "domain_model"
	"errors"
	"slices"
	"strings"
)

type MockDB struct {
	ActorsTable []dm.Actors
	FilmsTable  []dm.Films
	AFRelTable  []dm.ActorsFilmsRelations
	UsersTable  []dm.Users
}

func NewMoMockDB() *MockDB {
	return &MockDB{
		ActorsTable: []dm.Actors{},
		FilmsTable:  []dm.Films{},
		AFRelTable:  []dm.ActorsFilmsRelations{},
		UsersTable:  []dm.Users{},
	}
}

func (me *MockDB) AddActor(actor dm.Actors) error {
	idx := slices.IndexFunc(me.ActorsTable, func(a dm.Actors) bool {
		return a == actor
	})
	if idx != -1 {
		return errors.New("actor \"" + actor.Name + "\" already exists")
	}
	me.ActorsTable = append(me.ActorsTable, actor)
	return nil
}

func (me *MockDB) UpdateActor(name string, actor dm.Actors) error {
	idx := slices.IndexFunc(me.ActorsTable, func(a dm.Actors) bool {
		return a.Name == name
	})
	if idx == -1 {
		return errors.New("actor \"" + name + "\" does not exists")
	}
	me.ActorsTable[idx].Merge(actor)
	return nil
}

func (me *MockDB) DeleteActor(name string) error {
	idx := slices.IndexFunc(me.ActorsTable, func(a dm.Actors) bool {
		return a.Name == name
	})
	if idx == -1 {
		return errors.New("actor \"" + name + "\" does not exists")
	}
	me.ActorsTable = append(me.ActorsTable[:idx], me.ActorsTable[idx+1:]...)
	return nil
}

func (me *MockDB) GetActor(name string) (dm.Actors, error) {
	idx := slices.IndexFunc(me.ActorsTable, func(a dm.Actors) bool {
		return a.Name == name
	})
	if idx == -1 {
		return dm.Actors{}, errors.New("actor \"" + name + "\" does not exists")
	}
	return me.ActorsTable[idx], nil
}

func (me *MockDB) GetActorStarredFilms(name string) ([]dm.Films, error) {
	idx := slices.IndexFunc(me.ActorsTable, func(a dm.Actors) bool {
		return a.Name == name
	})
	if idx == -1 {
		return nil, errors.New("actor \"" + name + "\" does not exists")
	}
	actor_id_to_find := me.ActorsTable[idx].ID
	films_actor_starred := []dm.Films{}
	for _, v := range me.AFRelTable {
		if v.ActorID == actor_id_to_find {
			idx := slices.IndexFunc(me.FilmsTable, func(a dm.Films) bool {
				return a.ID == v.FilmID
			})
			if idx != -1 {
				films_actor_starred = append(films_actor_starred, me.FilmsTable[idx])
			}
		}
	}
	return films_actor_starred, nil
}

func (me *MockDB) AddFilm(film dm.Films) error {
	me.FilmsTable = append(me.FilmsTable, film)
	return nil
}

func (me *MockDB) AddFilmActors(film string, actors []dm.Actors) error {
	idx := slices.IndexFunc(me.FilmsTable, func(a dm.Films) bool {
		return a.Name == film
	})
	if idx == -1 {
		return errors.New("film \"" + film + "\" does not exists")
	}
	for _, v := range actors {
		me.AFRelTable = append(
			me.AFRelTable,
			dm.ActorsFilmsRelations{
				ActorID: v.ID,
				FilmID:  me.FilmsTable[idx].ID,
			},
		)
		me.ActorsTable = append(me.ActorsTable, actors...)
	}
	return nil
}

func (me *MockDB) DeleteFilmActor(film string, name string) error {
	idx := slices.IndexFunc(me.FilmsTable, func(a dm.Films) bool {
		return a.Name == film
	})
	if idx == -1 {
		return errors.New("film \"" + film + "\" does not exists")
	}

	idx = slices.IndexFunc(me.AFRelTable, func(a dm.ActorsFilmsRelations) bool {
		return a.FilmID == me.FilmsTable[idx].ID && me.FilmsTable[idx].Name == name
	})
	if idx != -1 {
		me.AFRelTable = append(me.AFRelTable[:idx], me.AFRelTable[idx+1:]...)
	}

	idx = slices.IndexFunc(me.ActorsTable, func(a dm.Actors) bool {
		return a.Name == name
	})
	if idx == -1 {
		me.ActorsTable = append(me.ActorsTable[:idx], me.ActorsTable[idx+1:]...)
	}
	return nil
}

func (me *MockDB) UpdateFilm(name string, film dm.Films) error {
	idx := slices.IndexFunc(me.FilmsTable, func(a dm.Films) bool {
		return a.Name == name
	})
	if idx == -1 {
		return errors.New("film \"" + name + "\" does not exists")
	}
	me.FilmsTable[idx].Merge(film)
	return nil
}

func (me *MockDB) DeleteFilm(name string) error {
	for i, v := range me.FilmsTable {
		if v.Name == name && len(me.FilmsTable) > i {
			for j, rel := range me.AFRelTable {
				if rel.FilmID == v.ID && len(me.AFRelTable) > j {
					me.AFRelTable = append(me.AFRelTable[:j], me.AFRelTable[j+1:]...)
				}
			}
			me.FilmsTable = append(me.FilmsTable[:i], me.FilmsTable[i+1:]...)
		}
	}
	return nil
}

func (me *MockDB) GetAllFilms(order dm.SortOrder, column dm.SortColumn) ([]dm.Films, error) {
	return me.FilmsTable, nil
}

func (me *MockDB) GetFilmsByNameSegment(segment string) ([]dm.Films, error) {
	films := []dm.Films{}
	for _, v := range me.FilmsTable {
		if strings.Contains(v.Name, segment) {
			films = append(films, v)
		}
	}
	if len(films) == 0 {
		return nil, errors.New("film with name segment \"" + segment + "\" does not exists")
	}
	return films, nil
}

func (me *MockDB) GetFilmsByActorNameSegment(segment string) ([]dm.Films, error) {
	films := []dm.Films{}
	for _, actor := range me.ActorsTable {
		if strings.Contains(actor.Name, segment) {
			idx := slices.IndexFunc(me.AFRelTable, func(a dm.ActorsFilmsRelations) bool {
				return a.ActorID == actor.ID
			})
			if idx == -1 {
				continue
			}
			film_id_to_find := me.AFRelTable[idx].FilmID
			idx = slices.IndexFunc(me.FilmsTable, func(a dm.Films) bool {
				return a.ID == film_id_to_find
			})
			if idx != -1 {
				films = append(films, me.FilmsTable[idx])
			}
		}
	}
	return films, nil
}

func (me *MockDB) GetFilmCast(name string) ([]dm.Actors, error) {
	film_idx := slices.IndexFunc(me.FilmsTable, func(a dm.Films) bool {
		return a.Name == name
	})
	if film_idx == -1 {
		return nil, errors.New("film \"" + name + "\" does not exists")
	}
	actors := []dm.Actors{}
	for _, v := range me.AFRelTable {
		if v.FilmID == me.FilmsTable[film_idx].ID {
			actor_idx := slices.IndexFunc(me.ActorsTable, func(a dm.Actors) bool {
				return a.ID == v.FilmID
			})
			if actor_idx != -1 {
				actors = append(actors, me.ActorsTable[actor_idx])
			}
		}
	}
	return actors, nil
}

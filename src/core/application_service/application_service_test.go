package application_service_test

import (
	as "application_service"
	"db"
	dm "domain_model"
	rep "repository"
	"testing"
	"time"
)

var service as.Filmoteka = as.NewFilmoteka(db.NewMoMockDB())
var mock_birthday int64 = time.Now().Unix()

type ActorTestDTO struct {
	ID       int64
	Name     string
	Gender   string
	Birthday int64
}

func (me *ActorTestDTO) MapToDomain() (dm.Actors, error) {
	return dm.Actors{
		ID:       me.ID,
		Name:     me.Name,
		Gender:   me.Gender,
		Birthday: me.Birthday,
	}, nil
}

func (me *ActorTestDTO) MapFromDomain(actor dm.Actors) {
	me.ID = actor.ID
	me.Name = actor.Name
	me.Gender = actor.Gender
	me.Birthday = actor.Birthday
}

type FilmTestDTO struct {
	ID          int64
	Name        string
	Description string
	ReseaseDate int64
	Rating      int
}

func (me *FilmTestDTO) MapToDomain() (dm.Films, error) {
	return dm.Films{
		ID:          me.ID,
		Name:        me.Name,
		Description: me.Description,
		ReseaseDate: me.ReseaseDate,
		Rating:      me.Rating,
	}, nil
}

func (me *FilmTestDTO) MapFromDomain(film dm.Films) {
	me.ID = film.ID
	me.Name = film.Name
	me.Description = film.Description
	me.ReseaseDate = film.ReseaseDate
	me.Rating = film.Rating
}

func TestAddActor(t *testing.T) {
	err := service.AddActor(&ActorTestDTO{
		ID:       1,
		Name:     "Vasilisa",
		Gender:   "Female",
		Birthday: mock_birthday,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetActor(t *testing.T) {
	actor, err := service.GetActor("Vasilisa")
	if err != nil {
		t.Fatal(err.Error())
	}

	if actor.ID != 1 ||
		actor.Name != "Vasilisa" ||
		actor.Gender != "Female" ||
		actor.Birthday != mock_birthday {
		t.Fatal("have got wrong actor")
	}
}

func TestUpdateActor(t *testing.T) {
	updated_actor := &ActorTestDTO{
		ID:       1,
		Name:     "Vasilisk",
		Gender:   "Male",
		Birthday: mock_birthday,
	}
	err := service.UpdateActor("Vasilisa", updated_actor)
	if err != nil {
		t.Fatal(err.Error())

	}
	actor, err := service.GetActor("Vasilisk")
	if err != nil {
		t.Fatal(err.Error())

	}
	if actor.ID != 1 ||
		actor.Name != "Vasilisk" ||
		actor.Gender != "Male" ||
		actor.Birthday != mock_birthday {
		t.Fatal("failed updating actor")
	}
}

func TestDeleteActor(t *testing.T) {
	err := service.DeleteActor("Vasilisk")
	if err != nil {
		t.Fatal(err.Error())

	}
	_, err = service.GetActor("Vasilisk")
	if err == nil {
		t.Fatal("deleted but should not")

	}
}

func TestAddFilm(t *testing.T) {
	release, _ := time.Parse(time.DateOnly, "1989")
	err := service.AddFilm(
		&FilmTestDTO{
			ID:          1,
			Name:        "Kiki's Delivery Service",
			Description: "Kiki and Zizi are going toward the adventures",
			ReseaseDate: release.Unix(),
			Rating:      10,
		},
	)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestAddFilmActors(t *testing.T) {
	actor_1 := ActorTestDTO{
		ID:       1,
		Name:     "Kiki",
		Gender:   "Female",
		Birthday: mock_birthday,
	}
	actor_2 := ActorTestDTO{
		ID:       2,
		Name:     "Zizi",
		Gender:   "Male",
		Birthday: mock_birthday,
	}
	var actors []rep.IDTO[dm.Actors]
	actors = append(actors, &actor_1, &actor_2)
	err := service.AddFilmActors("Kiki's Delivery Service", actors)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetActorStarredFilms(t *testing.T) {
	film, err := service.GetActorStarredFilms("Kiki")
	if err != nil {
		t.Fatal(err.Error())
	}

	release, _ := time.Parse(time.DateOnly, "1989")
	if film[0].Name != "Kiki's Delivery Service" ||
		film[0].ID != 1 ||
		film[0].Description != "Kiki and Zizi are going toward the adventures" ||
		film[0].ReseaseDate != release.Unix() ||
		film[0].Rating != 10 {
		t.Fatal("got wrong actor starred film")
	}
}

func TestDeleteFilmActor(t *testing.T) {
	err := service.DeleteFilmActor("Kiki's Delivery Service", "Zizi")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestUpdateFilm(t *testing.T) {
	release, _ := time.Parse(time.DateOnly, "1989")
	err := service.UpdateFilm(
		"Kiki's Delivery Service",
		&FilmTestDTO{
			ID:          1,
			Name:        "Kiki's Delivery Service",
			Description: "Young witch Kiki and her cat Zizi are going toward the adventures",
			ReseaseDate: release.Unix(),
			Rating:      10,
		})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetAllFilms(t *testing.T) {
	films, err := service.GetAllFilms(dm.SortOrder(dm.Descending), dm.SortColumn(dm.Rating))
	if err != nil {
		t.Fatal(err.Error())
	}
	release, _ := time.Parse(time.DateOnly, "1989")
	if len(films) == 0 ||
		films[0].Name != "Kiki's Delivery Service" ||
		films[0].ID != 1 ||
		films[0].Description != "Young witch Kiki and her cat Zizi are going toward the adventures" ||
		films[0].ReseaseDate != release.Unix() ||
		films[0].Rating != 10 {
		t.Fatal("the film left is wrong or there is no films")
	}
}

func TestGetFilmsByNameSegment(t *testing.T) {
	films, err := service.GetFilmsByNameSegment("Delivery Ser")
	if err != nil {
		t.Fatal(err.Error())
	}
	release, _ := time.Parse(time.DateOnly, "1989")
	if len(films) == 0 ||
		films[0].Name != "Kiki's Delivery Service" ||
		films[0].ID != 1 ||
		films[0].Description != "Young witch Kiki and her cat Zizi are going toward the adventures" ||
		films[0].ReseaseDate != release.Unix() ||
		films[0].Rating != 10 {
		t.Fatal("the film found is wrong or there is no films")
	}
}

func TestGetFilmsByActorNameSegment(t *testing.T) {
	films, err := service.GetFilmsByActorNameSegment("Ki")
	if err != nil {
		t.Fatal(err.Error())
	}
	release, _ := time.Parse(time.DateOnly, "1989")
	if len(films) == 0 ||
		films[0].Name != "Kiki's Delivery Service" ||
		films[0].ID != 1 ||
		films[0].Description != "Young witch Kiki and her cat Zizi are going toward the adventures" ||
		films[0].ReseaseDate != release.Unix() ||
		films[0].Rating != 10 {
		t.Fatal("the film found is wrong or there is no films")
	}
}

func TestGetFilmCast(t *testing.T) {
	actors, err := service.GetFilmCast("Kiki's Delivery Service")
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(actors) == 0 ||
		actors[0].ID != 1 ||
		actors[0].Name != "Kiki" ||
		actors[0].Gender != "Female" {
		t.Fatal("no film actors found or found wrong actor", actors)
	}
}

func TestDeleteFilm(t *testing.T) {
	err := service.DeleteFilm("Kiki's Delivery Service")
	if err != nil {
		t.Fatal(err.Error())
	}
	films, err := service.GetAllFilms(dm.Ascending, dm.Name)
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(films) > 0 {
		t.Fatal("film have not been deleted")
	}
}

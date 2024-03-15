package domain_model

type Actors struct {
	ID       int64
	Name     string
	Gender   string
	Birthday int64
}

type Films struct {
	ID          int64
	Name        string
	Description string
	ReseaseDate int64
	Rating      int
}

type ActorsFilmsRelations struct {
	ActorID int64
	FilmID  int64
}

type Users struct {
	ID       int64
	Role     string
	Login    string
	Password string
}

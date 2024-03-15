package domain_model

type SortOrder int

const (
	Descending SortOrder = 0
	Ascending  SortOrder = 1
)

type SortColumn int

const (
	Name        SortColumn = 0
	Rating      SortColumn = 1
	ReleaseDate SortColumn = 2
)

type Actors struct {
	ID       int64
	Name     string
	Gender   string
	Birthday int64
}

func (me *Actors) Merge(src Actors) {
	me.ID = GetIfNotOr(src.ID, -1, me.ID)
	me.Name = GetIfNotOr(src.Name, "", me.Name)
	me.Gender = GetIfNotOr(src.Gender, "", me.Gender)
	me.Birthday = GetIfNotOr(src.Birthday, -1, me.Birthday)
}

type Films struct {
	ID          int64
	Name        string
	Description string
	ReseaseDate int64
	Rating      int
}

func (me *Films) Merge(src Films) {
	me.ID = GetIfNotOr(src.ID, -1, me.ID)
	me.Name = GetIfNotOr(src.Name, "", me.Name)
	me.Description = GetIfNotOr(src.Description, "", me.Description)
	me.ReseaseDate = GetIfNotOr(src.ReseaseDate, -1, me.ReseaseDate)
	me.Rating = GetIfNotOr(src.Rating, -1, me.Rating)
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
	JWT      string
	Expiring int64
}

func GetIfNotOr[T comparable](new_v, cmp_v, alt_v T) T {
	if new_v != cmp_v {
		return new_v
	}
	return alt_v
}

package repository

import (
	dm "domain_model"
)

type IDTO[Domain any] interface {
	MapToDomain() (Domain, error)
	MapFromDomain(Domain)
}

type IRepository interface {
	AddActor(actor IDTO[dm.Actors]) error
	// UpdateActor(actor IDTO[dm.Actors]) error
}

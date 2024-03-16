package db

import (
	dm "domain_model"
	rep "repository"
)

type PostgredDB struct {
}

func NewPostgredDB() *PostgredDB {
	return &PostgredDB{}
}

func (me *PostgredDB) AddActor(actor rep.IDTO[dm.Actors]) error {
	return nil
}

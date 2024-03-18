package db

import (
	dm "domain_model"
	rep "repository"
)

type PostgresDB struct {
}

func NewPostgresDB() *PostgresDB {
	return &PostgresDB{}
}

func (me *PostgresDB) AddActor(actor rep.IDTO[dm.Actors]) error {
	return nil
}

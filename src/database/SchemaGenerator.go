package database

import (
	"github.com/jmoiron/sqlx"
)

type SchemaGenerator struct {
	Connection *sqlx.DB
}

func (s *SchemaGenerator) GenerateSchema(entities ...string) {
	for _, val := range entities {
		s.Connection.MustExec(val)
	}
}

package repository

import (
	"github.com/jmoiron/sqlx"
)

const (
	fileTable = "files"
)

type Repositories struct {
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{}
}

package postgres

import "database/sql"

type Repository struct {
	DB *sql.DB
}

func New(db *sql.DB) Repository {
	return Repository{
		DB: db,
	}
}

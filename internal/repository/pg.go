package repository

import "database/sql"

type PgRepo struct {
	DB *sql.DB
}

func (pg *PgRepo) doSomething() {}

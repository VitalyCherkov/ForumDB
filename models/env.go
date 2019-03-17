package models

import "github.com/jmoiron/sqlx"

// Env describes outer environment for route
type Env struct {
	DB *sqlx.DB
}

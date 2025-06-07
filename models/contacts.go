package models

import (
	"github.com/jackc/pgx/v5"
)

type Contact struct {
	Name		string
	Number 		uint64
	Favorite    bool
	Notes 		string
}


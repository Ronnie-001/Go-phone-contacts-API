package models

import (
	"encoding/json"
)

type Contact struct {
	Name		string		`json:"name"`
	Number 		uint64 		`json:"number"`
	Favorite    bool		`json:"favorite"`
	Notes 		string		`json:"notes"`
}


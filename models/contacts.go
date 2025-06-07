package models

import (
)

type Contact struct {
	Name		string     `json:"name"`
	Number 		string     `json:"number"`
	Favorite    bool       `json:"favorite"`
	Notes 		string     `json:"notes"`
}


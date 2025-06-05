package models

import (
	"github.com/jackc/pgx/v5"
)

type Contact struct {
	Name		string		`json:"name"`
	Number 		uint64 		`json:"number"`
	Favorite    bool		`json:"favorite"`
	Notes 		string		`json:"notes"`
}

func (c *Contact) SaveContact() (*Contact, error) {
	return c, err
}

func (c *Contact) GetContact() (*Contact, error) {
	return c, err
}  

func (c *Contact) DeleteContact() (*Contact, error) {
	return c, err
}

func (c *Contact) FavoriteContact() (*Contact, error) {
	return c, err
}

func (c *Contact) AddNote() (*Contact, error) {
	return c, err
}


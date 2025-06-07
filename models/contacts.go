package models

type Contact struct {
	Id 			int 	   `json:"contact_id"`
	Name		string     `json:"name"`
	Number 		string     `json:"number"`
	Favorite    bool       `json:"favorite"`
	Notes 		string     `json:"notes"`
}


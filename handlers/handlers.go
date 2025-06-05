package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go-phone-contacts-api/models"
)

type Home struct{}

func (h *Home) GoHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Phone contacts API"))
} 

type ContactsHandler struct{}

func (ch *ContactsHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	var model models.Contact
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		fmt.Fprintf( os.Stderr, "Error decoding HTTP request body: %v", err)
	}
}

func (ch *ContactsHandler) AddContact(w http.ResponseWriter, r *http.Request) {
}

func (ch *ContactsHandler) RemoveContact(w http.ResponseWriter, r *http.Request) {
}

func (ch *ContactsHandler) FavoriteContact(w http.ResponseWriter, r *http.Request) {
}

func (ch *ContactsHandler) AddNotes(w http.ResponseWriter, r *http.Request) {
}

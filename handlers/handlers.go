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

func (ch *ContactsHandler) AddContact(w http.ResponseWriter, r *http.Request) {
	var newContact models.Contact
	
	err := json.NewDecoder(r.Body).Decode(&newContact)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding HTTP request body: %v", err)
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newContact)
}

func (ch *ContactsHandler) GetContact(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (ch *ContactsHandler) RemoveContact(w http.ResponseWriter, r *http.Request) {
	
	w.WriteHeader(http.StatusOK)
}

func (ch *ContactsHandler) FavoriteContact(w http.ResponseWriter, r *http.Request) {
	
	w.WriteHeader(http.StatusOK)
}

func (ch *ContactsHandler) AddNotes(w http.ResponseWriter, r *http.Request) {
	
	w.WriteHeader(http.StatusOK)
}

package handlers

import (
	"context"
	"go-phone-contacts-api/models"

	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Home struct{}

func (h *Home) GoHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
} 

type ContactsHandler struct{
	db *pgxpool.Pool
}

func (ch *ContactsHandler) AddContact(w http.ResponseWriter, r *http.Request) {
	var newContact models.Contact
			
	// Decode the JSON body that has been recieved
	err := json.NewDecoder(r.Body).Decode(&newContact)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding HTTP request body: %v", err)
	}

	// Create & execute SQL statement to add Contact into the database
	query := "INSERT INTO contactsdb (name, number, favorite, notes) VALUES ($1, $2, $3)"
	_, err = ch.db.Exec(context.Background(), query, newContact.Name, newContact.Number, newContact.Favorite, newContact.Notes)
	if err != nil {
		fmt.Fprintf( os.Stderr, "Error adding contact into db: %v", err)
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

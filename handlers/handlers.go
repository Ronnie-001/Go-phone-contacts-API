package handlers

import (
	"context"
	"go-phone-contacts-api/models"
	"strings"

	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Home struct{}

func (h *Home) GoHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home handler"))
} 

type ContactsHandler struct {
	Db *pgxpool.Pool
}

func (ch *ContactsHandler) Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("testing"))
}

func (ch *ContactsHandler) AddContact(w http.ResponseWriter, r *http.Request) {
	var newContact models.Contact
			
	// Decode the JSON body that has been recieved
	err := json.NewDecoder(r.Body).Decode(&newContact)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding HTTP request body: %v", err)
	}

	// Create & execute SQL statement to add Contact into the database.
	query := "INSERT INTO contacts (name, number, favorite, notes) VALUES ($1, $2, $3, $4)"
	_, err = ch.Db.Exec(context.Background(), query, newContact.Name, newContact.Number, newContact.Favorite, newContact.Notes)
	if err != nil {
		fmt.Fprintf( os.Stderr, "Error adding contact into database: %v", err)
	}	
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newContact)
}

func (ch *ContactsHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact
	query := "SELECT * FROM contacts WHERE contact_id=$1"
	
	// Grab the id from the URL path by trimming it's prefix.
	id := strings.TrimPrefix(r.URL.Path, "/api/v1/getContact/") 

	err := ch.Db.QueryRow(context.Background(), query, id).Scan(&contact.Id, &contact.Name, &contact.Number, &contact.Favorite, &contact.Notes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching contact from database: %v", err)
	}	
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contact)
}

func (ch *ContactsHandler) RemoveContact(w http.ResponseWriter, r *http.Request) {
	
	// delete a contact from the database
	query := "DELETE FROM contacts WHERE contact_id=$1"

	id := strings.TrimPrefix(r.URL.Path, "/api/v1/removeContact/") 
	
	_, err := ch.Db.Exec(context.Background(), query, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error removing contact from the database: %v", err)
	}

	w.WriteHeader(http.StatusOK)
}

func (ch *ContactsHandler) FavoriteContact(w http.ResponseWriter, r *http.Request) {
	
	// Favorite a number. 
	query := "UPDATE contacts SET favorite=true WHERE contact_id=$1"

	id := strings.TrimPrefix(r.URL.Path, "/api/v1/favoriteContact/") 
	
	_, err := ch.Db.Exec(context.Background(), query, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error upddating favorite contact in database: %v", err)
	}

	w.WriteHeader(http.StatusOK)
}

func (ch *ContactsHandler) UnfavoriteContact(w http.ResponseWriter, r *http.Request) {
	
	// Unfavorite a number.
	query := "UPDATE contacts SET favorite=false WHERE contact_id=$1"

	id := strings.TrimPrefix(r.URL.Path, "/api/v1/unfavoriteContact/") 
	
	_, err := ch.Db.Exec(context.Background(), query, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error upddating favorite contact in database: %v", err)
	}

	w.WriteHeader(http.StatusOK)
}

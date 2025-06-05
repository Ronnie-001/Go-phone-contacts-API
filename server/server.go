package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-phone-contacts-api/handlers"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv" // used to load .env file
)

type Server struct {
	mux *http.ServeMux
}

func CreateServer() *Server {
	s := &Server{
		mux: http.NewServeMux(),
	}	

	return s
}

func (serv *Server) defineRoutes() {	
	homeHandler := &handlers.Home{}

	http.HandleFunc("GET /", homeHandler.GoHome)
	
	contactsHandler := &handlers.ContactsHandler{}

	http.HandleFunc("GET /api/v1/getContact", contactsHandler.GetContact)

	http.HandleFunc("POST /api/v1/addContact", contactsHandler.AddContact)
	
	http.HandleFunc("DELETE /api/v1/removeContact", contactsHandler.RemoveContact)

	http.HandleFunc("PUT /api/v1/favoriteContact", contactsHandler.FavoriteContact)

	http.HandleFunc("PUT /api/v1/addNote", contactsHandler.AddNotes)
}

func (serv *Server) connectToDatabase() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}	

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL")) // add the database URL here.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to the database: %v\n", err)	
		os.Exit(1)
	}

	defer conn.Close(context.Background())
}

func StartServer() {
	server := CreateServer()
	server.defineRoutes()
	http.ListenAndServe(":8080", server.mux)
}

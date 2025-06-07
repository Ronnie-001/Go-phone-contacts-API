package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-phone-contacts-api/handlers"

	"github.com/jackc/pgx/v5/pgxpool"	
	"github.com/joho/godotenv" // used to load .env file
)

type Server struct {
	mux *http.ServeMux
	db *pgxpool.Pool
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
	
	// Defining the routes for the API
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
	
	// Creating a connection pool
	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to the database: %v\n", err)	
		os.Exit(1)
	}
	
	// Adding connection pool to the Server struct
	serv.db = db
}

func StartServer() {
	server := CreateServer()
	server.defineRoutes()
	http.ListenAndServe(":8080", server.mux)
}

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

// Group dependencies
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
	serv.mux.HandleFunc("GET /", homeHandler.GoHome)
	
	// Defining the routes for the API
	// FIXME: fix init of pgxpool driver in this handler
	contactsHandler := &handlers.ContactsHandler{db: serv.db}
	serv.mux.HandleFunc("GET /api/v1/getContact", contactsHandler.GetContact)

	serv.mux.HandleFunc("POST /api/v1/addContact", contactsHandler.AddContact)
	
	serv.mux.HandleFunc("DELETE /api/v1/removeContact", contactsHandler.RemoveContact)

	serv.mux.HandleFunc("PUT /api/v1/favoriteContact", contactsHandler.FavoriteContact)

	serv.mux.HandleFunc("PUT /api/v1/addNote", contactsHandler.AddNotes)
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

func (serv *Server) pingDB(ctx context.Context) error {
	return serv.db.Ping(ctx)
}

func StartServer() {
	server := CreateServer()
	server.connectToDatabase()
	server.defineRoutes()
	server.pingDB(context.Background())
	http.ListenAndServe(":8080", server.mux)
}

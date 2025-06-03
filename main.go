package main 

import (
	"net/http"
	"go-phone-contacts-api/handlers"
)

func main() {
	
	// creating empty HTTP router
	mux := http.NewServeMux()
	
	mux.Handle("/", &handlers.Home{})

	// listen on localhost
	http.ListenAndServe(":8080", mux)
}

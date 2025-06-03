package handlers

import (
	"net/http"
)

type Home struct{}

func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Simple Bookstore API"))
} 



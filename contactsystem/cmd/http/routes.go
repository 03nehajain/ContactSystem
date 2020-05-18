package main

import (
	"log"
	"net/http"

	model "bitbucket.org/exotel/exotel_code/contactsystem/internal/core/contract"
	"github.com/gorilla/mux"
)

func AddRoutes() {
	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/api/contacts", model.GetContacts).Methods("GET")
	r.HandleFunc("/api/contact/{name}", model.GetContact).Methods("GET")
	r.HandleFunc("/api/contacts/{email}", model.GetContact).Methods("GET")
	r.HandleFunc("/api/contacts", model.CreateContact).Methods("POST")
	r.HandleFunc("/api/contact/{email}", model.UpdateContact).Methods("PUT")
	r.HandleFunc("/api/contacts/{email}", model.DeleteContact).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}

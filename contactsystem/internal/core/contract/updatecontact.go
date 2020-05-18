package contract

import (
	"encoding/json"
	"net/http"

	dbconn "bitbucket.org/exotel/exotel_code/contactsystem/internal/connmanager"
	mdl "bitbucket.org/exotel/exotel_code/contactsystem/internal/models"
	"github.com/gorilla/mux"
)

// Update contact
func UpdateContact(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var contact mdl.Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)

	db := dbconn.ConnectDatabse()

	stmt, err := db.Prepare("UPDATE contact set name =? ,mobile =? where email =?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(contact.Name, contact.Mobile, params["email"])
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(contact)

}

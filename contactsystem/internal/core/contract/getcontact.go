package contract

import (
	"encoding/json"
	"net/http"

	dbconn "bitbucket.org/exotel/exotel_code/contactsystem/internal/connmanager"
	mdl "bitbucket.org/exotel/exotel_code/contactsystem/internal/models"
	"github.com/gorilla/mux"
)

// Get single contact
func GetContact(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var contacts []mdl.Contact
	params := mux.Vars(r)
	db := dbconn.ConnectDatabse()

	rows, err := db.Query(`SELECT name,email,mobile FROM contact where email =? or name =?`, params["email"], params["name"])
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		contact := mdl.Contact{}
		err = rows.Scan(&contact.Name, &contact.Email, &contact.Mobile)
		if err != nil {
			panic(err)
		}
		contacts = append(contacts, contact)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(contacts)
}

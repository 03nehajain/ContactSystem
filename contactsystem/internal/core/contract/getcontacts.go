package contract

import (
	"encoding/json"
	"net/http"

	dbconn "bitbucket.org/exotel/exotel_code/contactsystem/internal/connmanager"
	mdl "bitbucket.org/exotel/exotel_code/contactsystem/internal/models"
)

// Get all contacts
func GetContacts(w http.ResponseWriter, r *http.Request) {
	db := dbconn.ConnectDatabse()
	rows, err := db.Query(`SELECT name,email,mobile FROM contact limit 10`)
	if err != nil {
		panic(err)
	}
	//  Init contacts var as a slice Contact struct
	var contacts []mdl.Contact
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)

}

package contract

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	dbconn "bitbucket.org/exotel/exotel_code/contactsystem/internal/connmanager"
	mdl "bitbucket.org/exotel/exotel_code/contactsystem/internal/models"
)

// Add new Contact
func CreateContact(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var contact mdl.Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)

	contact.ID = strconv.Itoa(rand.Intn(100000000))
	db := dbconn.ConnectDatabse()

	stmt, err := db.Prepare("INSERT INTO contact(id,name,email,mobile) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(contact.ID, contact.Name, contact.Email, contact.Mobile)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(contact)
}

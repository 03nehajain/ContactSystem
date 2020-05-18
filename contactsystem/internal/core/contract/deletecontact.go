package contract

import (
	"net/http"

	dbconn "bitbucket.org/exotel/exotel_code/contactsystem/internal/connmanager"
	"github.com/gorilla/mux"
)

// Delete contact
func DeleteContact(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db := dbconn.ConnectDatabse()

	_, err := db.Exec("DELETE FROM contact where email =?", params["email"])
	if err != nil {
		panic(err)
	}

}

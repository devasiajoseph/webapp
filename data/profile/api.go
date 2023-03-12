package profile

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var apiObj = "profile"

func SaveApi(w http.ResponseWriter, r *http.Request) {

}

func ListApi(w http.ResponseWriter, r *http.Request) {

}

func DeleteApi(w http.ResponseWriter, r *http.Request) {

}

func AddRoutes(r *mux.Router) {
	r.HandleFunc("/api/"+apiObj, SaveApi).Methods("POST")
}

// Start initializes bitcoin based functions
func Start(r *mux.Router) {
	log.Println("Starting " + apiObj + " api")
	AddRoutes(r)
}

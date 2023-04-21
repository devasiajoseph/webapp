package location

import (
	"log"
	"net/http"

	"github.com/devasiajoseph/webapp/libs/api"
	"github.com/gorilla/mux"
)

var apiObj = "location"

func getCountriesApi(w http.ResponseWriter, r *http.Request) {
	cl, err := GetCountryList()
	if err != nil {
		api.ServerError(w)
	}

	api.ObjectResponse(w, cl)
}
func AddRoutes(r *mux.Router) {
	r.HandleFunc("/api/"+apiObj+"/countries", getCountriesApi).Methods("GET")
}

// Start initializes this package functions
func Start(r *mux.Router) {
	log.Println("Starting " + apiObj + " api")
	AddRoutes(r)
}

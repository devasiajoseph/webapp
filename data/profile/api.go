package profile

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devasiajoseph/webapp/api"
	"github.com/devasiajoseph/webapp/uauth"
	"github.com/gorilla/mux"
)

var apiObj = "profile"

func (obj *Object) hasAuth(w http.ResponseWriter, r *http.Request) bool {
	ua, err := uauth.GetAuthenticatedUser(r)
	if err != nil {
		log.Println(err)
		api.ServerError(w)
		return false
	}

	obj.ProfileID = api.PostInt(r, "profile_id")
	auth, err := obj.IsManager(ua)

	if err != nil {
		log.Println(err)
		api.ServerError(w)
		return false
	}
	if !auth {
		api.AuthError(w)
		return false
	}
	return true
}

func saveApi(w http.ResponseWriter, r *http.Request) {
	obj := Object{}
	if !obj.hasAuth(w, r) {
		return
	}

}

func listApi(w http.ResponseWriter, r *http.Request) {

}

func deleteApi(w http.ResponseWriter, r *http.Request) {

}

func getApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	obj := Object{}
	if !obj.hasAuth(w, r) {
		return
	}

}

func AddRoutes(r *mux.Router) {
	r.HandleFunc("/api/"+apiObj, saveApi).Methods("POST")
	r.HandleFunc("/api/"+apiObj, getApi).Methods("GET")
}

// Start initializes bitcoin based functions
func Start(r *mux.Router) {
	log.Println("Starting " + apiObj + " api")
	AddRoutes(r)
}

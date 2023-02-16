package bitcoin

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devasiajoseph/webapp/api"
	"github.com/devasiajoseph/webapp/bitcoin/blockcypher"
	"github.com/gorilla/mux"
)

func GetBalanceApi(w http.ResponseWriter, r *http.Request) {
	balance, err := blockcypher.GetBalance(api.QueryParam(r, "addr"))
	if err != nil {
		log.Println(err)
		api.ServerError(w)
		return
	}
	fmt.Println(balance)
	api.ObjectResponse(w, balance)
}

func AddRoutes(r *mux.Router) {
	r.HandleFunc("/api/bitcoin/balance", GetBalanceApi).Methods("GET")
}

// Start initializes bitcoin based functions
func Start(r *mux.Router) {
	log.Println("Starting bitcoin api")
	AddRoutes(r)
}

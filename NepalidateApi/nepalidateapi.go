package nepalidateapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NepaliApi() {
	router := mux.NewRouter()
	router.HandleFunc("/nepalidate", GetNepaliDate).Methods("GET")
	http.ListenAndServe(":8000", router)

}

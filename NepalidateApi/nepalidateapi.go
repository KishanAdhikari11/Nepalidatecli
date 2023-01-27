package nepalidateapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	scrapper "nepalidate/Nepalidatescrapper"

	"github.com/gorilla/mux"
)

func NepaliDateApi() {
	router := mux.NewRouter()
	router.HandleFunc("/nepalidate", GetNepaliDate).Methods("GET")
	fmt.Println("Server started on port 8000")
	http.ListenAndServe(":8000", router)

}

func GetNepaliDate(w http.ResponseWriter, r *http.Request) {
	date := scrapper.Scrape()
	fmt.Println(date)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&date)

}

package dinoapi

import (
	"dino/databaselayer"
	"net/http"

	"github.com/gorilla/mux"
)

//dino API =>
// HTTP GET for search /api/dinos/nickname/rex , or search for /api/dinos/type/velociraptor
// HTTP POST to add or edit /api/dinos/add or /api/dinos/edit
// HTTP POST to /api/dinos/edit/rex, rex will be edited using the data obtained from the http json body of the request
// The HTTP POST request comes with a json body which hosts the data to be added or used to the edit

func RunApi(endpoint string, db databaselayer.DinoDBHandler) error {
	//endpoint example: localhost:8080
	r := mux.NewRouter()
	RunAPIOnRouter(r, db)
	return http.ListenAndServe(endpoint, r)
}

func RunAPIOnRouter(r *mux.Router, db databaselayer.DinoDBHandler) {
	handler := newDinoRESTAPIHandler(db)

	apirouter := r.PathPrefix("/api/dinos").Subrouter()

	apirouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.searchHandler)
	apirouter.Methods("POST").PathPrefix("/{Operation}").HandlerFunc(handler.editsHandler)
}

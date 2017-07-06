package dinoapi

import (
	"dino/databaselayer"
	"fmt"
	"log"
	"net/http"

	"strings"

	"encoding/json"

	"github.com/gorilla/mux"
)

type DinoRESTAPIHandler struct {
	dbhandler databaselayer.DinoDBHandler
}

func newDinoRESTAPIHandler(db databaselayer.DinoDBHandler) *DinoRESTAPIHandler {
	return &DinoRESTAPIHandler{
		dbhandler: db,
	}
}

func (handler *DinoRESTAPIHandler) searchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `No search criteria found, you can either search by nickname via /api/dinos/nickname/rex , or
						to search by type via /api/dinos/type/velociraptor`)
		return
	}

	searchkey, ok := vars["search"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `No search criteria found, you can either search by nickname via /api/dinos/nickname/rex , or
						to search by type via /api/dinos/type/velociraptor`)
		return
	}
	var animal databaselayer.Animal
	var animals []databaselayer.Animal
	var err error
	switch strings.ToLower(criteria) {
	case "nickname":
		animal, err = handler.dbhandler.GetDynoByNickname(searchkey)
	case "type":
		animals, err = handler.dbhandler.GetDynosByType(searchkey)
	}
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error occured while querying animals %v ", err)
		return
	}
	if len(animals) > 0 {
		json.NewEncoder(w).Encode(animals)
		return
	}
	json.NewEncoder(w).Encode(animal)
}

func (handler *DinoRESTAPIHandler) editsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operation, ok := vars["Operation"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `Operation was not provided, you can either use /api/dinos/add to add a new animal, or
						/api/dinos/edit/rex to edit an existing animal data with nickname rex`)
		return
	}
	var animal databaselayer.Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Could not docode the request body to json %v", err)
		return
	}
	switch strings.ToLower(operation) {
	case "add":
		err = handler.dbhandler.AddAnimal(animal)
	case "edit":
		nickname := r.RequestURI[len("api/dinos/edit/")+1:]
		log.Println("edit requested for nickname", nickname)
		err = handler.dbhandler.UpdateAnimal(animal, nickname)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error occured while processing request %v", err)
	}
}

package dinowebportal

import (
	"dino/databaselayer"
	"dino/dinowebportal/dinoapi"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//RunWebPortal starts running the dino web portal on address addr
func RunWebPortal(dbtype uint8, addr, dbconnection, frontend string) error {
	r := mux.NewRouter()

	initializeAPI(dbtype, dbconnection, r)
	fileserver := http.FileServer(http.Dir(frontend))
	r.PathPrefix("/").Handler(fileserver)
	return http.ListenAndServe(addr, r)
}

func initializeAPI(dbtype uint8, dbconnection string, r *mux.Router) error {
	db, err := databaselayer.GetDatabaseHandler(dbtype, dbconnection)
	if err != nil {
		return err
	}
	dinoapi.RunAPIOnRouter(r, db)
	return nil
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Dino web portal %s", r.RemoteAddr)
}

package handler

import (
	"github.com/gorilla/mux"
	"go-run/storage"
	"net/http"
)

type application struct {
	DB storage.Service
}

func GetHandler(dbService storage.Service) http.Handler {
	app := application{DB: dbService}

	r := mux.NewRouter()
	// add default handlers like MethodNotAllowed, Resource doesn't exist etc
	r.HandleFunc("/user", app.createUser).Methods(http.MethodPost)
	return r
}

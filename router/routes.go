package router

import (
	// "command-line-arguments/Users/poulav/Documents/My_GoLang_Code/15GoMongo/controller/controller.go"
	"github.com/poulav/gomongodb/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/create", controller.CreateMyMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deletemovies", controller.DeleteAllMovies).Methods("DELETE")

}

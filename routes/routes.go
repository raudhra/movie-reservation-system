package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := http.NewServesMux()

	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{ID}", getMovieById).Methods("GET")
	r.HandleFunc("/movie", addMovie).Methods("POST")
	r.HandleFunc("/movie/{ID}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{ID}", deleteMovie).Methods("DELETE")

	log.Println("Server Starting On Port :8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

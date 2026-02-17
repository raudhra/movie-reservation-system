package routes

import (
	"net/http"
)

func main() {
	r := http.NewServesMux()

	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{ID}", getMovieById).Methods("GET")
	r.HandleFunc("/movie", addMovie).Methods("POST")
	r.HandleFunc("/movie/{ID}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{ID}", deleteMovie).Methods("DELETE")
}

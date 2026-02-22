package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/raudhra/movie-reservation-system/models"
)

var NewMovie models.Movie

func getAllMovies(w http.ResponseWriter r *http.Request) {
	newMovie := models.GetAllMovies()
	res, _:= json.Marshal(newMovie)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func getMovie(w http.ResponseWriter, r *http.Request) {

}

func addMovie(w http.ResponseWriter, r *http.Request) {

}

func updateMovie(w http.ResponseWriter,r *http.Request) {

}

func deleteMovie(w http.ResponseWriter,r *http.Request) {

}

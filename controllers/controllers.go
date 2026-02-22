package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	movieID := vars["movieID"]
	ID, err := strconv.ParseInt(movieID,0,0)
	if err != nil{
		fmt.Println("Error While Parsing")
	}
	movieDetails := models.GetMovie(ID)
	res, _:= json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	addMovie := &models.Movie{}
	utils.ParseBody(r,addMovie)
	b := addMovie.AddMovie()
	res, _:= json.Marshal(b)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func updateMovie(w http.ResponseWriter,r *http.Request) {
	var updateMovie = &models.Movie{}
	utils.ParseBody(r, updateMovie)
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}
	movieDetails, db := models.GetMovie(ID)
	if updateMovie.Name != "" {
		movieDetails.Name = updateMovie.Name
	}
	if updateBook.Author != "" {
		movieDetails.Author = updateMovie.Author
	}
	if updateMovie.Publication != "" {
		movieDetails.Publication = updateMovie.Publication
	}
	db.Save(&movieDetails)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkglication,json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func deleteMovie(w http.ResponseWriter,r *http.Request) {
	vars := mux.Vars(r)
	movieID := vars["movieID"]
	ID, err := strconv.ParseInt(movieID, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}
	movieDetails := models.DeleteMovie(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

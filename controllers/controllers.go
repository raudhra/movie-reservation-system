package controllers

import (
	"net/http"
	"strconv"

	"github.com/raudhra/movie-reservation-system/models"
)

func getAllMovies(c *gin.Context) {
	movies := models.GetAllMovies()
	c.JSON(http.StatusOK, movies)
}

func getMovie(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to convert ID to int",
		})
		return
	}
	movie := models.GetMovie(uint(intId))
	c.JSON(http.StatusOK, movie)
}

func addMovie(c *gin.Context) {
	model := models.Movie{}
	if err := c.ShouldBindJSON(&model); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Error parsing movie to json",
		})
		return
	}
	model.AddMovie()
	c.JSON(http.StatusCreated, model)
}

func updateMovie(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to convert ID to int",
		})
		return
	}
	model := models.Movie{}
	if err := c.ShouldBindJSON(&model); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Error parsing movie to json",
		})
		return
	}
	updated := models.UpdateMovie(uint(intId), model)
	c.JSON(http.StatusOK, updated)
}

func deleteMovie(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to convert ID to int",
		})
		return
	}
	movie := models.DeleteMovie(uint(intId))
	c.JSON(http.StatusOK, movie)
}

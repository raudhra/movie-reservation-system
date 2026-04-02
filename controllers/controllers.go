package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raudhra/movie-reservation-system/models"
)

func GetAllMovies(c *gin.Context) {
	movies := models.GetAllMovies()
	c.JSON(http.StatusOK, movies)
}

func GetMovie(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to convert ID to int",
		})
		return
	}
	movie, _ := models.GetMovie(uint(intId))
	c.JSON(http.StatusOK, movie)
}

func AddMovie(c *gin.Context) {
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

func UpdateMovie(c *gin.Context) {
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

func DeleteMovie(c *gin.Context) {
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

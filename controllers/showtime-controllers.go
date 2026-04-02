package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raudhra/movie-reservation-system/models"
)

func GetAllShowtimes(c *gin.Context) {
	showtimes := models.GetAllShowtimes()
	c.JSON(http.StatusOK, showtimes)
}

func GetShowtime(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to convert ID to int",
		})
		return
	}
	showtime, _ := models.GetShowtime(uint(intId))
	c.JSON(http.StatusOK, showtime)
}

func AddShowtime(c *gin.Context) {
	model := models.Showtimes{}
	if err := c.ShouldBindJSON(&model); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Error parsing showtime to json",
		})
		return
	}
	condition := models.CheckOverlap(model.MovieID, model.StartTime)
	if condition == true {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "There already is an existing showtime in this timeframe",
		})
		return
	}
	model.AddShowtime()
	c.JSON(http.StatusCreated, model)
}

func UpdateShowtime(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to convert ID to int",
		})
		return
	}
	model := models.Showtimes{}
	if err := c.ShouldBindJSON(&model); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Error parsing showtime to json",
		})
		return
	}
	updated := models.UpdateShowtime(uint(intId), model)
	c.JSON(http.StatusOK, updated)
}

func DeleteShowtime(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to convert ID to int",
		})
		return
	}
	showtime := models.DeleteShowtime(uint(intId))
	c.JSON(http.StatusOK, showtime)
}

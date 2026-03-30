package controllers

import (
	"net/http"
	"strconv"

	"github.com/raudhra/movie-reservation-system/models"
)

func getAllReservation(c *gin.Context) {
	reservations := models.GetAllReservation()
	c.JSON(http.StatusOK, reservations)
}

func getUserReservations(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to convert ID to int",
		})
		return
	}
	reservations := models.GetUserReservation(uint(intId))
	c.JSON(http.StatusOK, reservations)
}

func createReservation(c *gin.Context) {
	model := models.Reservation{}
	if err := c.ShouldBindJSON(&model); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Error parsing reservation to json",
		})
		return
	}
	model.CreateReservation()
	c.JSON(http.StatusCreated, model)
}

func pdateReservation(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to convert ID to int",
		})
		return
	}
	model := models.Reservation{}
	if err := c.ShouldBindJSON(&model); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Error parsing reservation to json",
		})
		return
	}
	updated := models.UpdateReservation(uint(intId), model)
	c.JSON(http.StatusOK, updated)
}

func cancelReservation(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to convert ID to int",
		})
		return
	}
	reservation := models.CancelReservation(uint(intId))
	c.JSON(http.StatusOK, reservation)
}

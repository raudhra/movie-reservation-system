package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/raudhra/movie-reservation-system/models"
)

func getAllReservation(c *gin.Context) {
	reservations := models.GetAllReservation()
	c.JSON(http.StatusOK, reservations)
}

func getUserReservations(c *gin.Context) {
	userId, _ := c.Get("userID")
	id := userId.(uint)
	reservations := models.GetUserReservation(uint(id))
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
	check := model.ShowtimeID
	showtimes, _ := models.GetShowtime(check)
	if showtimes.AvailableSeats == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "There are no available seats",
		})
		return
	}
	model.CreateReservation()
	c.JSON(http.StatusCreated, model)
}

func updateReservation(c *gin.Context) {
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
	reservation, _ := models.GetReservation(uint(intId))
	showtime, _ := models.GetShowtime(reservation.ShowtimeID)
	if showtime.StartTime.Before(time.Now()) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "Cannot cancel when showtime has started",
		})
		return
	}
	cancellation := models.CancelReservation(uint(intId))
	c.JSON(http.StatusOK, cancellation)
}

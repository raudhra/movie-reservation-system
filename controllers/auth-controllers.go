package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raudhra/movie-reservation-system/authentication"
	"github.com/raudhra/movie-reservation-system/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Error while binding User",
		})
		return
	}
	existing := models.GetUserByEmail(user.Email)
	if existing.ID != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "User is already registered",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Error while hashing password",
		})
		return
	}
	user.Password = string(hash)
	user.CreateUser()
	c.JSON(http.StatusCreated, gin.H{"name": user.Name, "email": user.Email, "role": user.Role})

}

func Login(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Error While Binding",
		})
		return
	}
	existing := models.GetUserByEmail(user.Email)
	if existing.ID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "User is not registered",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(existing.Password), []byte(user.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Wrong Password",
		})
		return
	}
	token, err := authentication.GenerateToken(existing.ID, existing.Email, string(existing.Role))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": "Unable to generate JWT token",
		})
		return
	}
	c.JSON(http.StatusOK, token)

}

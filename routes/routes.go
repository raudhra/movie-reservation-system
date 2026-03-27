package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/raudhra/movie-reservation-system/middleware"
)

func routes() {
	router := gin.Default()
	router.POST("/register", register)
	router.POST("/login", login)
	signingKey := os.Getenv("JWT_SECRET")
	//user routes
	user := router.Group("/user")
	user.Use(middleware.AuthMiddleware(signingKey))
	{
		user.GET("/movies", getAllMovies)
		user.GET("/movies/:id", getMovie)
		user.GET("/showtimes", getshowtime)
		user.GET("/showtimes/:id", getshowtime)
		user.POST("/reservations", reservation)
	}
	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware(signingKey))
	{
		admin.GET("/movies", getAllMovies)
		admin.GET("/movies/:id", getMovie)
		admin.GET("/showtimes", getAllShowtimes)
		admin.GET("/showtimes/:id", getShowtime)
		admin.POST("/reservations", reservation)
		admin.POST("/movies", addMovie)
		admin.POST("/showtimes", addShowtime)
		admin.PUT("/movies/:id", updateMovie)
		admin.DELETE("/movies/:id", deleteMovie)
		admin.PUT("/showtime/:id", updateShowtime)
		admin.DELETE("/showtime/:id", deleteShowtime)
	}
	return *gin.Engine
}

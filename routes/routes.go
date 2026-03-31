package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/raudhra/movie-reservation-system/controllers"
	"github.com/raudhra/movie-reservation-system/middleware"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	signingKey := os.Getenv("JWT_SECRET")
	//user routes
	user := router.Group("/user")
	user.Use(middleware.AuthMiddleware(signingKey))
	{
		user.GET("/movies", controllers.GetAllMovies)
		user.GET("/movies/:id", controllers.GetMovie)
		user.GET("/showtimes", controllers.GetAllShowtime)
		user.GET("/showtimes/:id", controllers.GetShowtime)
		user.POST("/reservations", controllers.CreateReservation)
		user.GET("/reservations/:id", controllers.GetUserReservations)
		user.DELETE("/reservations/:id", controllers.CancelReservation)
	}
	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware(signingKey))
	admin.Use(middleware.AdminOnly())
	{
		admin.GET("/movies", controllers.GetAllMovies)
		admin.POST("/movies", controllers.AddMovie)
		admin.GET("/movies/:id", controllers.GetMovie)
		admin.PUT("/movies/:id", controllers.UpdateMovie)
		admin.DELETE("/movies/:id", controllers.DeleteMovie)
		admin.GET("/showtimes", controllers.GetAllShowtimes)
		admin.GET("/showtimes/:id", controllers.GetShowtime)
		admin.POST("/showtimes", controllers.AddShowtime)
		admin.PUT("/showtime/:id", controllers.UpdateShowtime)
		admin.DELETE("/showtimes/:id", controllers.DeleteShowtime)
		admin.POST("/reservations", controllers.CreateReservation)
		admin.GET("/reservations", controllers.GetAllReservation)
		admin.GET("/reservations/:id", controllers.GetUserReservations)
		admin.PUT("/reservations/:id", controllers.UpdateReservation)
		admin.DELETE("/reservations/:id", controllers.CancelReservation)

	}
	return router
}

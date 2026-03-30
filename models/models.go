package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/raudhra/movie-reservation-system/config"
)

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleGuest UserRole = "guest"
)

type Status string

const (
	ConfirmedStatus Status = "confirmed"
	CancelledStatus Status = "cancelled"
)

type Movie struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"unique;not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Genre       string    `gorm:"not null" json:"genre"`
	Duration    int       `gorm:"not null" json:"duration"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type User struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"unique;not null" json:"name"`
	Email      string    `gorm:"unique;not null" json:"email"`
	Password   string    `gorm:"not null" json:"password"`
	SignupTime time.Time `json:"signuptime"`
	Role       UserRole  `gorm:"not null;default:guest" json:"role"`
}

type Showtimes struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	MovieID        uint      `gorm:"not null" json:"movieid"`
	StartTime      time.Time `gorm:"not null" json:"starttime"`
	TotalSeats     int       `gorm:"not null" json:"totalseats"`
	AvailableSeats int       `json:"availableseats"`
}

type Reservation struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	ShowtimeID uint   `gorm:"uniqueIndex:composites;not null" json:"showtimeid"`
	UserID     uint   `gorm:"not null" json:"userid"`
	SeatNumber int    `gorm:"uniqueIndex:composites" json:"seatnumber"`
	Status     Status `gorm:"not null" json:"status"`
}

var db *gorm.DB

func (m *Movie) ValidateMovie() error {
	if m.Title == "" {
		return errors.New("Title cannot be empty")
	}

	if m.Description == "" {
		return errors.New("Description cannot be empty")
	}

	if m.Genre == "" {
		return errors.New("Genre cannot be empty")
	}

	return nil
}

func init() {
	config.Connect()
	db = config.getDB()
	db.AutoMigrate(
		&Movie{},
		&User{},
		&Showtimes{},
		&Reservation{},
	)
}

func GetAllMovies() []Movie {
	var Movies []Movie
	db.Find(&Movies)
	return Movies
}

func GetMovie(ID uint) (*Movie, *gorm.DB) {
	var getMovie Movie
	db := db.Where("ID=?", ID).First(&getMovie, ID)
	return &getMovie, db
}

func DeleteMovie(ID uint) Movie {
	var movie Movie
	db.First(&movie, ID)
	db.Delete(&movie)
	return movie
}

func (m *Movie) AddMovie() *Movie {
	db.Create(&m)
	return m
}

func UpdateMovie(ID uint, updatedMovie Movie) Movie {
	var movie Movie
	db.First(&movie, ID)
	db.Model(&movie).Updates(updatedMovie)
	return movie
}

func GetUserByEmail(Email string) *User {
	var user User
	db.Where("email=?", Email).First(&user)
	return &user
}

func (u *User) CreateUser() *User {
	db.Create(&u)
	return u
}

func GetAllShowtimes() []Showtimes {
	var Showtime []Showtimes
	db.Find(&Showtime)
	return Showtime
}

func GetShowtime(ID uint) (*Showtimes, *gorm.DB) {
	var getShowtime Showtimes
	db := db.Where("ID=?", ID).First(&getShowtime, ID)
	return &getShowtime, db
}

func (s *Showtimes) AddShowtime() *Showtimes {
	db.Create(&s)
	return s
}

func UpdateShowtime(ID uint, updatedShowtime Showtimes) *Showtimes {
	var showtime Showtimes
	db.First(&showtime, ID)
	db.Model(&showtime).Updates(updatedShowtime)
	return &showtime
}

func DeleteShowtime(ID uint) *Showtimes {
	var showtime Showtimes
	db.First(&showtime, ID)
	db.Delete(&showtime)
	return &showtime
}

func CheckOverlap(movieID uint, startTime time.Time) bool {
	var count int64
	db.Model(&Showtimes{}).
		Joins("join movies on movies.id = showtimes.movie_id").
		Where("showtimes.movie_id = ? AND ? >= showtimes.start_time AND ? < showtimes.start_time + (movies.duration * interval '1 minute')", movieID, startTime, startTime).
		Count(&count)
	return count > 0
}

func GetAllReservation() []Reservation {
	var Reservations []Reservation
	db.Find(&Reservations)
	return Reservations
}

func GetUserReservation(ID uint) (*Reservation, *gorm.DB) {
	var getUserReservation Reservation
	db := db.Where("ID=?", ID).First(&getUserReservation, ID)
	return &getUserReservation, db
}

func (r *Reservation) CreateReservation() *Reservation {
	db.Create(&r)
	return r
}

func UpdateReservation(ID uint, updatedReservation Reservation) *Reservation {
	var reservation Reservation
	db.First(&reservation, ID)
	db.Model(&reservation).Updates(updatedReservation)
	return &reservation
}

func CancelReservation(ID uint) *Reservation {
	var reservation Reservation
	db.First(&reservation, ID)
	db.Delete(&reservation)
	return &reservation
}

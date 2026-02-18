package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Movie struct {
	ID          uint      `gorn:"primaryKey" json:"id"`
	Title       string    `gorm:"unique;not null" json:"title"`
	Description string    `gorm:"unique;not null" json:"description"`
	Genre       string    `gorm:"unique;not null" json:"genre"`
	Showtime    string    `gorm:"unique;not null" json:"showtime"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (m *Movie) ValidateMovie() error {
	if m.Name == "" {
		return errors.New("Title cannot be empty")
	}

	if m.Description == "" {
		return errors.New("Description cannot be empty")
	}

	if m.Genre == "" {
		return errors.New("Genre cannot be empty")
	}

	if m.Showtime == "" {
		return errors.New("Showtime cannot be empty")
	}
}

func init() {
	config.Connect()
	db = config.getDB()
	db.AutoMigrate(&Book{})
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
	db.NewRecord(m)
	db.Create(&m)
	return m
}

func UpdateMovie(ID uint) Movie {
	var movie Movie
	db.First(&movie, ID)
	db.Delete(&movie)
	db.NewRecord(movie)
	db.Create(&movie)
	return movie
}

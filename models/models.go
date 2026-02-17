package models

import (
	"errors"
	"time"
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

package models

import "time"

type Post struct {
	ID        uint `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string `json:"title" gorm:"unique;not null;default:null"`
	Body      string `json:"body" gorm:"not null;default:null"`
}

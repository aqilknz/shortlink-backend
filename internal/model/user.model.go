package model

import "time"

type User struct {
	Id           int       `json:"id" db:"id"`
	ProfilePhoto *string   `json:"profilePhoto" db:"profile_photo"`
	FullName     *string   `json:"fullName" db:"fullname"`
	Email        string    `json:"email" db:"email"`
	Password     string    `json:"-" db:"password"`
	CreatedAt    time.Time `json:"created_at"`
}

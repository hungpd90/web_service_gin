package models

import "time"

type User struct {
	ID                  uint      `gorm:"primaryKey"`
	Email               string    `json:"email" binding:"required,email"`
	EmailVerified       string    `json:"email_verified" gorm:"default:false"`
	PhoneNumber         string    `json:"phone_number"`
	PhoneNumberVerified string    `json:"phone_number_verified" gorm:"default:false"`
	Password            string    `json:"password"`
	Otp                 string    `json:"otp"`
	FirstName           string    `json:"first_name" validate:"required"`
	LastName            string    `json:"last_name"`
	Dob                 time.Time `json:"dob"`
	Gender              string    `json:"gender"`
	ProfilePicture      string    `json:"profile_picture"`
	LastLoggedIn        time.Time `gorm:"default:null" json:"last_logged_in,omitempty"`
	IsDeleted           bool      `gorm:"default:false" json:"is_deleted,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

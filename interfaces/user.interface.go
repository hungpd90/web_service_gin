package interfaces

import "time"

type GetUserResponse struct {
	ID                  *uint      `json:"id,omitempty" gorm:"primaryKey"`
	Email               *string    `json:"email,omitempty" binding:"required,email"`
	EmailVerified       *string    `json:"email_verified,omitempty"`
	PhoneNumber         *string    `json:"phone_number,omitempty"`
	PhoneNumberVerified *string    `json:"phone_number_verified,omitempty"`
	Password            *string    `json:"password,omitempty"`
	Otp                 *string    `json:"otp,omitempty"`
	FirstName           *string    `json:"first_name,omitempty" validate:"required"`
	LastName            *string    `json:"last_name,omitempty"`
	Dob                 *time.Time `json:"dob,omitempty"`
	Gender              *string    `json:"gender,omitempty"`
	ProfilePicture      *string    `json:"profile_picture,omitempty"`
	LastLoggedIn        *time.Time `gorm:"default:null" json:"last_logged_in,omitempty"`
	IsDeleted           *bool      `gorm:"default:false" json:"is_deleted,omitempty"`
	CreatedAt           *time.Time `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}

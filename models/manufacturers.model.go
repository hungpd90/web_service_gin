package models

import "time"

type Manufacturer struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Origin    string    `json:"origin"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

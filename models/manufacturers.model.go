package models

import "time"

type Manufacturer struct {
	ID        string
	Name      string
	Origin    string
	IsDeleted bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

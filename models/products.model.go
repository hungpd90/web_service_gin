package models

import (
	"time"

	"github.com/lib/pq"
)

type Product struct {
	ID             uint `gorm:"primaryKey"`
	SKU            string
	Name           string
	Description    string
	Images         string
	Rating         float64
	Price          float64
	Quantity       uint
	SoldAmount     uint
	ManufacturerId pq.StringArray `gorm:"type:string[]"`
	Manufacturer   Manufacturer
	IsDeleted      bool `gorm:"default:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

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
	Images         pq.StringArray `gorm:"type:text[]"`
	Rating         float64
	Price          float64
	Quantity       uint
	SoldAmount     uint
	ManufacturerId uint
	Manufacturer   Manufacturer `gorm:"foreignKey:manufacturer_id"`
	IsDeleted      bool         `gorm:"default:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

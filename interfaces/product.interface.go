package interfaces

import (
	"time"
	"web-service-gin/models"

	"github.com/lib/pq"
)

type GetProductsResponse struct {
	ID             *uint                `json:"id,omitempty" gorm:"primaryKey"`
	SKU            *string              `json:"sku,omitempty"`
	Name           *string              `json:"name,omitempty"`
	Description    *string              `json:"description,omitempty"`
	Images         *pq.StringArray      `json:"images,omitempty" gorm:"type:string[]"`
	Rating         *float64             `json:"rating,omitempty"`
	Price          *float64             `json:"price,omitempty"`
	Quantity       *uint                `json:"quantity,omitempty"`
	SoldAmount     *uint                `json:"sold_amount,omitempty"`
	ManufacturerId *string              `json:"manufacturer_id,omitempty"`
	Manufacturer   *models.Manufacturer `json:"omitempty"`
	IsDeleted      *bool                `json:"is_deleted,omitempty" gorm:"default:false"`
	CreatedAt      *time.Time           `json:"created_at,omitempty"`
	UpdatedAt      *time.Time           `json:"updated_at,omitempty"`
}

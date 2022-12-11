package validates

import "github.com/lib/pq"

type CreateProductBody struct {
	SKU            string         `json:"sku" binding:"required"`
	Name           string         `json:"name" binding:"required"`
	Description    string         `json:"description" binding:"required"`
	Images         pq.StringArray `gorm:"type:string[]" json:"images"`
	Price          float64        `json:"price" binding:"required"`
	Quantity       uint           `json:"quantity" binding:"required"`
	ManufacturerId uint           `json:"manufacturer_id" binding:"required"`
}

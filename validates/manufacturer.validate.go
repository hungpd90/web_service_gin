package validates

type CreateManufacturerBody struct {
	Name   string `json:"name" binding:"required"`
	Origin string `json:"origin" binding:"required"`
}

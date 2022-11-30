package validates

type SignUpBody struct {
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	Dob         string `json:"dob" binding:"required"`
}

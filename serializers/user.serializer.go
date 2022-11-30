package serializers

import (
	"web-service-gin/interfaces"
	"web-service-gin/models"
)

func GetUserSerializer(user *models.User) interfaces.GetUserResponse {
	var response interfaces.GetUserResponse
	response.Email = &user.Email
	response.EmailVerified = &user.EmailVerified
	response.PhoneNumber = &user.PhoneNumber
	response.PhoneNumberVerified = &user.PhoneNumberVerified
	response.FirstName = &user.FirstName
	response.LastName = &user.LastName
	response.Dob = &user.Dob
	response.Gender = &user.Gender
	response.ProfilePicture = &user.ProfilePicture
	response.LastLoggedIn = &user.LastLoggedIn
	response.CreatedAt = &user.CreatedAt
	response.UpdatedAt = &user.UpdatedAt
	return response
}

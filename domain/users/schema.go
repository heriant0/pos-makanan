package users

type UserRequest struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	ImageUrl    string `json:"image_url"`
}

type UserResponse struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	ImageUrl    string `json:"image_url"`
}

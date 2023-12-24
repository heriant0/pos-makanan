package merchants

type MerchantRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	ImageUrl    string `json:"image_url"`
}

type MerchantResponse struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	ImageUrl    string `json:"image_url"`
}

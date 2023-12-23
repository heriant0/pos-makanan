package merchants

type Merchant struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	City        string `json:"city"`
	ImageUrl    string `json:"imageUrl"`
}

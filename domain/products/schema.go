package products

type ProductResponse struct {
	Id          int     `json:"id"`
	CategoryId  int     `json:"categoryId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	ImageUrl    string  `json:"image_url"`
}

type ProductRequest struct {
	CategoryId  int     `json:"categoryId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	ImageUrl    string  `json:"image_url"`
}

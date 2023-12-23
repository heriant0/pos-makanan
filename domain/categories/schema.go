package categories

type CategoryResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// CreateAt    time.Time `json:"createdAt"`
	// UpdatedAt   time.Time `json:"updatedAt"`
}

package categories

import "time"

type Category struct {
	Id          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	CreateAt    time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}

func (c Category) parseToCategoryResponse() CategoryResponse {
	return CategoryResponse{
		Id:          c.Id,
		Name:        c.Name,
		Description: c.Description,
	}
}

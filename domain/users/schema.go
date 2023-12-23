package users

import "time"

type UserRequest struct {
	Name        string    `jsong:"name"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	PhoneNumber string    `json:"phoneNumber"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	ImageUrl    string    `json:"imageUrl"`
}

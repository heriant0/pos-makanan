package users

import (
	"strings"
	"time"
)

type User struct {
	Id int `db:"id" json:"id"`
	Name        string    `db:"name" jsong:"name"`
	DateOfBirth time.Time `db:"date_of_birth" json:"dateOfBirth"`
	PhoneNumber string    `db:"phone_number" json:"phoneNumber"`
	Gender      string    `db:"gender" json:"gender"`
	Address     string    `db:"address" json:"address"`
	ImageUrl    string    `db:"image_url" json:"imageUrl"`
}

func requestBody(req UserRequest) (user User, err error) {
	user = User{
		Name:        req.Name,
		DateOfBirth: req.DateOfBirth,
		PhoneNumber: req.PhoneNumber,
		Gender:      req.Gender,
		Address:     req.Address,
		ImageUrl:    req.ImageUrl,
	}
	err = user.validate()
	return
}

func (u User) validate() error {
	if err := u.genderRequired(); err != nil {
		return err
	} else if err := u.invalidGender(); err != nil {
		return err
	} else if err := u.phoneNumberEmpty(); err != nil {
		return err
	} else if err := u.phoneNumberLength(); err != nil {
		return err
	} else if err := u.nameRequired(); err != nil {
		return err
	} else if err := u.addressRequired(); err != nil {
		return err
	} else if err := u.imageUrl(); err != nil {
		return err
	}

	return nil
}

func (u User) genderRequired() error {
	if u.Gender == "" {
		return GenderIsRequired
	}
	return nil
}

func (u User) invalidGender() error {
	gender := strings.ToLower(u.Gender)
	if gender != "male" && gender != "female" {
		return GenderIsInvalid
	}
	return nil
}

func (u User) phoneNumberEmpty() error {
	if u.PhoneNumber == "" {
		return PhoneNumberIsEmpty
	}
	return nil
}

func (u User) phoneNumberLength() error {
	if len(u.PhoneNumber) < 10 {
		return PhoneNumberLength
	}
	return nil
}

func (u User) nameRequired() error {
	if u.Name == "" {
		return NameIsRequired
	}
	return nil
}

func (u User) addressRequired() error {
	if u.Address == "" {
		return AddressIsRequired
	}
	return nil
}

// func (u User) dob() error {
// 	if u.B == "" {
// 		return DateOfBirthIsRequired
// 	}
// 	return nil
// }

// func (u User) x() error {
// 	if {
// 		return DateOfBirthIsInvalid
// 	}
// 	return nil
// }

func (u User) imageUrl() error {
	if u.ImageUrl == "" {
		return ImageUrlIsRequird
	}
	return nil
}

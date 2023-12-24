package merchants

type Merchant struct {
	Id          int    `db:"id", json:"id"`
	Name        string `db:"name" json:"name"`
	Address     string `db:"address" json:"address"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	City        string `db:"city" json:"city"`
	ImageUrl    string `db:"image_url" json:"image_url"`
}

func requestBody(req MerchantRequest) (merchat Merchant, err error) {
	merchat = Merchant{
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		City:        req.City,
		ImageUrl:    req.ImageUrl,
	}

	err = merchat.validate()
	return
}

func (m Merchant) validate() error {
	if err := m.nameIsRequired(); err != nil {
		return err
	} else if err := m.phoneNumberEmpty(); err != nil {
		return err
	} else if err := m.phoneNumberEmpty(); err != nil {
		return err
	} else if err := m.imageUrl(); err != nil {
		return err
	} else if err := m.cityRequired(); err != nil {
		return err
	}
	return nil
}

func (m Merchant) nameIsRequired() error {
	if m.Name == "" {
		return NameIsRequired
	}
	return nil
}

func (m Merchant) phoneNumberEmpty() error {
	if m.PhoneNumber == "" {
		return PhoneNumberIsEmpty
	}
	return nil
}

func (m Merchant) phoneNumberLength() error {
	if len(m.PhoneNumber) < 10 {
		return PhoneNumberLength
	}
	return nil
}

func (m Merchant) imageUrl() error {
	if m.ImageUrl == "" {
		return ImageUrlIsRequird
	}
	return nil
}

func (m Merchant) cityRequired() error {
	if m.City == "" {
		return CityIsRequired
	}
	return nil
}

package auth

import "net/mail"

type Auth struct {
	Id       int    `db:"id" json:"id"`
	Email    string `db:"email" json:"emaiL"`
	Password string `db:"password" json:"password"`
	Role     string `db:"role" json:"role"`
}

func requestBody(req AuthRequest) (auth Auth, err error) {
	auth = Auth{
		Email:    req.Email,
		Password: req.Password,
	}

	err = auth.validate()
	return
}

func (a Auth) validate() error {
	if err := a.emailRequire(); err != nil {
		return err
	} else if err := a.validateEmail(); err != nil {
		return err
	} else if err := a.passwordRequire(); err != nil {
		return err
	} else if err := a.passwordLenght(); err != nil {
		return err
	}

	return nil
}

func (a Auth) emailRequire() error {
	if a.Email == "" {
		return EmailIsRequired
	}
	return nil
}

func (a Auth) validateEmail() error {
	_, err := mail.ParseAddress(a.Email)
	if err != nil {
		return EmailIsInvalid
	}
	return nil
}

func (a Auth) passwordRequire() error {
	if a.Password == "" {
		return PasswordIsEmpty
	}
	return nil
}

func (a Auth) passwordLenght() error {
	if len(a.Password) < 6 {
		return PasswordLength
	}
	return nil
}

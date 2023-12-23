package utility

import (
	"github.com/heriant0/pos-makanan/internal/config"
	"golang.org/x/crypto/bcrypt"
)

var cfg config.Bcrypt

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), cfg.HasCost)

	return string(bytes)
}

func VerifyPassword(plainPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(plainPassword),
	)

	return err == nil
}

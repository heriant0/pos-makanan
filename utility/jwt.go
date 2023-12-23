package utility

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/heriant0/pos-makanan/internal/config"
)

var cfgToken config.Token
var secret = []byte(cfgToken.Secret)

func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   jwt.NewNumericDate(time.Now().Add(10 * time.Minute)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(token string) (jwt.MapClaims, error) {
	tokens, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	claims, ok := tokens.Claims.(jwt.MapClaims)
	if ok && tokens.Valid {
		return claims, nil
	}

	return jwt.MapClaims{}, fmt.Errorf("unable to extract claims")
}

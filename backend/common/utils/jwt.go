package utils

import (
	"errors"
	"my-posting-site/backend/common/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	expirationHours    = 24
	expirationMins     = 12
	ExpirationDuration = time.Duration(expirationHours*time.Hour + expirationMins*time.Minute)
)

var (
	secret = []byte("SomeSecret")
)

type Token struct {
	UserID  int       `json:"user_id,omitempty"`
	Email   string    `json:"email,omitempty"`
	TimeExp time.Time `json:"time_exp,omitempty"`
	jwt.StandardClaims
}

func GenerateJWT(user *models.User) string {
	expirationTime := time.Now().Local().Add(ExpirationDuration)

	tk := &Token{UserID: int(user.ID), Email: user.Email, TimeExp: expirationTime}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString(secret)

	return tokenString
}

func UnpackJWT(tokenHeader string) (*Token, error) {
	if tokenHeader == "" {
		return nil, errors.New("token is not present")
	}

	tk := &Token{}
	tokenValue := tokenHeader
	token, err := jwt.ParseWithClaims(tokenValue, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, errors.New("malformed authentication token")
	}

	diff := time.Until(tk.TimeExp)

	if diff < 0 {
		return nil, errors.New("token expired")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return tk, nil
}

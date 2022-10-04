package utils

import (
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

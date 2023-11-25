package helper

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = os.Getenv("JWT_KEY")
var SECRET_KEY = []byte(JwtKey)

func GenerateToken(userID int) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour).Unix() // expiration time as two hour from the current time.

	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = expirationTime // Set the expiration time.

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

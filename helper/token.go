package helper

import (
	"errors"
	"net/http"
	"os"
	"salmaqnsGH/sysnapsis-assessment/model/web"
	"strings"
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

func GetUserIDFromToken(writer http.ResponseWriter, req *http.Request) interface{} {
	authHeader := req.Header.Get("Authorization")

	if !strings.Contains(authHeader, "Bearer") {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		WriteToResponseBody(writer, webResponse)
		return webResponse
	}

	tokenString := ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}

	token, err := ValidateToken(tokenString)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		WriteToResponseBody(writer, webResponse)
		return webResponse
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		WriteToResponseBody(writer, webResponse)
		return webResponse
	}

	return claims["user_id"]
}

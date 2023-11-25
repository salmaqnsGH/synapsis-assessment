package middleware

import (
	"net/http"
	"salmaqnsGH/sysnapsis-assessment/helper"
	"salmaqnsGH/sysnapsis-assessment/model/web"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/api/v1/users/register" || req.URL.Path == "/api/v1/users/login" {
		middleware.Handler.ServeHTTP(writer, req)
		return
	}

	authHeader := req.Header.Get("Authorization")

	if !strings.Contains(authHeader, "Bearer") {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	tokenString := ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}

	token, err := helper.ValidateToken(tokenString)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	middleware.Handler.ServeHTTP(writer, req)
}

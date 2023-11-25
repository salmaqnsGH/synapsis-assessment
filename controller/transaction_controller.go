package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TransactionController interface {
	AddToCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

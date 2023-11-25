package controller

import (
	"net/http"
	"salmaqnsGH/sysnapsis-assessment/helper"
	"salmaqnsGH/sysnapsis-assessment/model/web"
	"salmaqnsGH/sysnapsis-assessment/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}

func (controller *TransactionControllerImpl) AddToCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionCreateRequest := web.CartCreateRequest{}
	helper.ReadFromRequestBody(request, &transactionCreateRequest)

	productIdParams := params.ByName("productId")
	productID, err := strconv.Atoi(productIdParams)
	helper.PanicIfError(err)

	userID := helper.GetUserIDFromToken(writer, request)

	transactionCreateRequest.UserID = int(userID.(float64))
	transactionCreateRequest.ProductID = productID

	productResponse := controller.TransactionService.AddToCart(request.Context(), transactionCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionControllerImpl) FindAllProductInCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userID := int(helper.GetUserIDFromToken(writer, request).(float64))

	productResponses := controller.TransactionService.FindAllProductInCart(request.Context(), userID)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userID := int(helper.GetUserIDFromToken(writer, request).(float64))

	productIdParams := params.ByName("productId")
	productID, err := strconv.Atoi(productIdParams)
	helper.PanicIfError(err)

	controller.TransactionService.Delete(request.Context(), productID, userID)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionControllerImpl) Checkout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartID := params.ByName("cartId")
	id, err := strconv.Atoi(cartID)
	helper.PanicIfError(err)

	transactionResponse := controller.TransactionService.Checkout(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

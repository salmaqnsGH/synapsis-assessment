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

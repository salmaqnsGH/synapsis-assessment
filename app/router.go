package app

import (
	"salmaqnsGH/sysnapsis-assessment/controller"
	"salmaqnsGH/sysnapsis-assessment/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/v1/categories", categoryController.FindAll)
	router.GET("/api/v1/categories/:categoryId", categoryController.FindByID)
	router.POST("/api/v1/categories", categoryController.Create)
	router.PUT("/api/v1/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/v1/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}

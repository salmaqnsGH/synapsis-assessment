package app

import (
	"salmaqnsGH/sysnapsis-assessment/controller"
	"salmaqnsGH/sysnapsis-assessment/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController, userController controller.UserController, productController controller.ProductController, transactionController controller.TransactionController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/cart/add/:productId", transactionController.AddToCart)
	router.DELETE("/api/v1/cart/:productId", transactionController.Delete)
	router.GET("/api/v1/cart/", transactionController.FindAllProductInCart)
	router.GET("/api/v1/cart/checkout/:cartId", transactionController.Checkout)

	router.GET("/api/v1/categories", categoryController.FindAll)
	router.GET("/api/v1/categories/:categoryId", categoryController.FindByID)
	router.POST("/api/v1/categories", categoryController.Create)
	router.PUT("/api/v1/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/v1/categories/:categoryId", categoryController.Delete)

	router.GET("/api/v1/products", productController.FindAll)
	router.GET("/api/v1/products/:productId", productController.FindByID)
	router.POST("/api/v1/products", productController.Create)
	router.PUT("/api/v1/products/:productId", productController.Update)
	router.DELETE("/api/v1/products/:productId", productController.Delete)

	router.POST("/api/v1/users/register", userController.Register)
	router.POST("/api/v1/users/login", userController.Login)
	router.POST("/api/v1/users/update", userController.Update)

	router.PanicHandler = exception.ErrorHandler

	return router
}

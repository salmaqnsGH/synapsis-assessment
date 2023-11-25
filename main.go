package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"salmaqnsGH/sysnapsis-assessment/app"
	"salmaqnsGH/sysnapsis-assessment/controller"
	"salmaqnsGH/sysnapsis-assessment/helper"
	"salmaqnsGH/sysnapsis-assessment/middleware"
	"salmaqnsGH/sysnapsis-assessment/repository"
	"salmaqnsGH/sysnapsis-assessment/service"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	db := app.NewDB()

	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	transactionRepository := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(transactionRepository, productRepository, db, validate)
	transactionController := controller.NewTransactionController(transactionService)

	router := app.NewRouter(categoryController, userController, productController, transactionController)

	address := fmt.Sprintf("%s:%s", host, port)

	server := http.Server{
		Addr:    address,
		Handler: middleware.NewAuthMiddleware(router),
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}

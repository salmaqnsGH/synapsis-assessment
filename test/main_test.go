package test

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"salmaqnsGH/sysnapsis-assessment/app"
	"salmaqnsGH/sysnapsis-assessment/controller"
	"salmaqnsGH/sysnapsis-assessment/repository"
	"salmaqnsGH/sysnapsis-assessment/service"
	"strconv"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

const (
	baseUrl    = "http://localhost"
	port       = 3000
	apiVersion = "api/v1"
)

func TestMain(m *testing.M) {
	// Run the tests
	code := m.Run()

	// You can do cleanup here if needed

	// Exit with the code from the test runner
	os.Exit(code)
}

func setupTestDB() *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	host := os.Getenv("TEST_DB_HOST")
	port := os.Getenv("TEST_DB_PORT")
	user := os.Getenv("TEST_DB_USER")
	password := os.Getenv("TEST_DB_PASSWORD")
	dbname := os.Getenv("TEST_DB_NAME")

	intPort, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Invalid port: %s", err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, intPort, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %s", err)
	}

	return db
}

func setupRouter(db *sql.DB) http.Handler {
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

	return router
}

func truncateCategory(db *sql.DB) {
	db.Exec("DELETE FROM categories;")
}

func truncateProduct(db *sql.DB) {
	db.Exec("DELETE FROM products;")
}

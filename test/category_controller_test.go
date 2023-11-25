package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"salmaqnsGH/sysnapsis-assessment/app"
	"salmaqnsGH/sysnapsis-assessment/controller"
	"salmaqnsGH/sysnapsis-assessment/model/domain"
	"salmaqnsGH/sysnapsis-assessment/repository"
	"salmaqnsGH/sysnapsis-assessment/service"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	baseUrl    = "http://localhost"
	port       = 3000
	apiVersion = "api/v1"
)

func setupTestDB() *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("TEST_DB_HOST")
	port := os.Getenv("TEST_DB_PORT")
	user := os.Getenv("TEST_DB_USER")
	password := os.Getenv("TEST_DB_PASSWORD")
	dbname := os.Getenv("TEST_DB_NAME")

	intPort, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, intPort, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
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

	router := app.NewRouter(categoryController, userController)

	return router
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE categories;")
}

func createCategory(db *sql.DB) domain.Category {
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name:        "category name",
		Description: "category description",
	})
	tx.Commit()

	return category
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	url := fmt.Sprintf("%s:%d/%s/categories",
		baseUrl, port, apiVersion)

	requestBody := strings.NewReader(`{"name":"new category name","description":"new category description"}`)
	request := httptest.NewRequest(http.MethodPost, url, requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "new category name", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "new category description", responseBody["data"].(map[string]interface{})["description"])
}
func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	url := fmt.Sprintf("%s:%d/%s/categories",
		baseUrl, port, apiVersion)

	requestBody := strings.NewReader(`{"name":"","description":""}`)
	request := httptest.NewRequest(http.MethodPost, url, requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)
	category := createCategory(db)

	url := fmt.Sprintf("%s:%d/%s/categories/%d",
		baseUrl, port, apiVersion, category.ID)

	requestBody := strings.NewReader(`{"name":"updated category name","description":"updated category description"}`)
	request := httptest.NewRequest(http.MethodPut, url, requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.ID, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "updated category name", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "updated category description", responseBody["data"].(map[string]interface{})["description"])
}
func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)
	category := createCategory(db)

	url := fmt.Sprintf("%s:%d/%s/categories/%d",
		baseUrl, port, apiVersion, category.ID)

	requestBody := strings.NewReader(`{"name":"","description":""}`)
	request := httptest.NewRequest(http.MethodPut, url, requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	truncateCategory(db)
	category := createCategory(db)

	url := fmt.Sprintf("%s:%d/%s/categories/%d",
		baseUrl, port, apiVersion, category.ID)

	request := httptest.NewRequest(http.MethodGet, url, nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.ID, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, category.Description, responseBody["data"].(map[string]interface{})["description"])
}
func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	url := fmt.Sprintf("%s:%d/%s/categories/%d",
		baseUrl, port, apiVersion, 404)

	request := httptest.NewRequest(http.MethodGet, url, nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)
	category := createCategory(db)

	url := fmt.Sprintf("%s:%d/%s/categories/%d",
		baseUrl, port, apiVersion, category.ID)

	request := httptest.NewRequest(http.MethodDelete, url, nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}
func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	url := fmt.Sprintf("%s:%d/%s/categories/%d",
		baseUrl, port, apiVersion, 404)

	request := httptest.NewRequest(http.MethodDelete, url, nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestGetListCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	truncateCategory(db)

	category1 := createCategory(db)
	category2 := createCategory(db)

	url := fmt.Sprintf("%s:%d/%s/categories",
		baseUrl, port, apiVersion)

	request := httptest.NewRequest(http.MethodGet, url, nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	var categories = responseBody["data"].([]interface{})

	categoryResponse1 := categories[0].(map[string]interface{})
	categoryResponse2 := categories[1].(map[string]interface{})

	assert.Equal(t, category1.ID, int(categoryResponse1["id"].(float64)))
	assert.Equal(t, category1.Name, categoryResponse1["name"])
	assert.Equal(t, category1.Description, categoryResponse1["description"])

	assert.Equal(t, category2.ID, int(categoryResponse2["id"].(float64)))
	assert.Equal(t, category2.Name, categoryResponse2["name"])
	assert.Equal(t, category2.Description, categoryResponse2["description"])
}

package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"salmaqnsGH/sysnapsis-assessment/model/domain"
	"salmaqnsGH/sysnapsis-assessment/repository"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	_ "github.com/lib/pq"
)

func createProduct(db *sql.DB, categoryID int) domain.Product {
	tx, _ := db.Begin()
	productRepository := repository.NewProductRepository()
	product := productRepository.Save(context.Background(), tx, domain.Product{
		Name:        "product name",
		Description: "product description",
		CategoryID:  categoryID,
	})
	tx.Commit()

	return product
}

func TestCreateProductSuccess(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)
	truncateCategory(db)
	category := createCategory(db)
	router := setupRouter(db)

	url := fmt.Sprintf("%s:%d/%s/products",
		baseUrl, port, apiVersion)

	req := fmt.Sprintf(`{"name":"new product name","description":"new product description", "category_id":%d}`, category.ID)
	requestBody := strings.NewReader(req)
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
	assert.Equal(t, category.ID, int(responseBody["data"].(map[string]interface{})["category_id"].(float64)))
	assert.Equal(t, "new product name", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "new product description", responseBody["data"].(map[string]interface{})["description"])
}

func TestCreateProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)
	truncateCategory(db)
	category := createCategory(db)
	router := setupRouter(db)

	url := fmt.Sprintf("%s:%d/%s/products",
		baseUrl, port, apiVersion)

	req := fmt.Sprintf(`{"name":"","description":"", "category_id":%d}`, category.ID)

	requestBody := strings.NewReader(req)
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

func TestUpdateProductSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	category := createCategory(db)
	product := createProduct(db, category.ID)
	router := setupRouter(db)

	url := fmt.Sprintf("%s:%d/%s/products/%d",
		baseUrl, port, apiVersion, product.ID)

	req := fmt.Sprintf(`{"name":"updated product name","description":"updated product description", "category_id":%d}`, category.ID)

	requestBody := strings.NewReader(req)
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
	assert.Equal(t, product.ID, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.ID, int(responseBody["data"].(map[string]interface{})["category_id"].(float64)))
	assert.Equal(t, "updated product name", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "updated product description", responseBody["data"].(map[string]interface{})["description"])
}

func TestUpdateProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)
	category := createCategory(db)
	product := createProduct(db, category.ID)
	router := setupRouter(db)

	url := fmt.Sprintf("%s:%d/%s/products/%d",
		baseUrl, port, apiVersion, product.ID)

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

func TestGetProductSuccess(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)
	category := createCategory(db)
	product := createProduct(db, category.ID)
	router := setupRouter(db)

	url := fmt.Sprintf("%s:%d/%s/products/%d",
		baseUrl, port, apiVersion, product.ID)

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
	assert.Equal(t, product.ID, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, product.CategoryID, int(responseBody["data"].(map[string]interface{})["category_id"].(float64)))
	assert.Equal(t, product.Name, responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, product.Description, responseBody["data"].(map[string]interface{})["description"])
}

func TestGetProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)

	router := setupRouter(db)

	url := fmt.Sprintf("%s:%d/%s/products/%d",
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

func TestDeleteProductSuccess(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)
	category := createCategory(db)
	product := createProduct(db, category.ID)
	router := setupRouter(db)

	url := fmt.Sprintf("%s:%d/%s/products/%d",
		baseUrl, port, apiVersion, product.ID)

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

func TestDeleteProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)
	router := setupRouter(db)

	url := fmt.Sprintf("%s:%d/%s/products/%d",
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

func TestGetListProductsSuccess(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)
	category := createCategory(db)
	product1 := createProduct(db, category.ID)
	product2 := createProduct(db, category.ID)
	router := setupRouter(db)

	url := fmt.Sprintf("%s:%d/%s/products",
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

	var products = responseBody["data"].([]interface{})

	productResponse1 := products[0].(map[string]interface{})
	productResponse2 := products[1].(map[string]interface{})

	assert.Equal(t, product1.ID, int(productResponse1["id"].(float64)))
	assert.Equal(t, product1.CategoryID, int(productResponse1["category_id"].(float64)))
	assert.Equal(t, product1.Name, productResponse1["name"])
	assert.Equal(t, product1.Description, productResponse1["description"])

	assert.Equal(t, product2.ID, int(productResponse2["id"].(float64)))
	assert.Equal(t, product2.CategoryID, int(productResponse2["category_id"].(float64)))
	assert.Equal(t, product2.Name, productResponse2["name"])
	assert.Equal(t, product2.Description, productResponse2["description"])
}

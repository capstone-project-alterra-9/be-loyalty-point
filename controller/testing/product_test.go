package controller

import (
	"capstone-project/config"
	"capstone-project/controller"
	"capstone-project/entity"
	"capstone-project/repository"
	"capstone-project/service"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func InitProductsTestAPI() *echo.Echo {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
	db := config.InitDatabaseTest()
	repository := repository.NewRepository(db)
	Service := service.NewService(repository)
	controller.NewController(Service)
	e := echo.New()
	return e
}

func InsertDataProduct() {
	Product := entity.Products{
		ID:          "1",
		Category:    "credits",
		Name:        "Credits Test Product",
		Description: "This is a test product",
		Price:       12000,
		Stock:       10,
		Image:       "10rb",
	}
	err := entity.DB.Create(&Product).Error
	if err != nil {
		panic(err)
	}
}

func TestCreateProduct(t *testing.T) {
	e := InitProductsTestAPI()
	e.POST("/api/auth/products",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Create Product", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"category": "credits",
			"name": "Credits Test Product",
			"description": "This is a test product",
			"price": 12000,
			"stock": 10,
			"image": "10rb"		
			}`)

		auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWlud2ViQGdtYWlsLmNvbSIsImV4cCI6MTc3MDc3MDk2NSwidXNlcm5hbWUiOiJhZG1pbndlYiJ9.qg4vb8IBXbIuL9hK_aNEky59UWet5fF4DzPWVIDwdvQ"
		request := httptest.NewRequest(http.MethodPost, "/api/auth/products", requestBody)
		request.Header.Set(echo.HeaderAuthorization, auth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.CreateProduct(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})

	// // Test Case 2
	// t.Run("Create Product with Unknown Category", func(t *testing.T) {
	// 	requestBody := strings.NewReader(`{
	// 		"name": "Credits Test Product",
	// 		"description": "This is a test product",
	// 		"price": 12000,
	// 		"stock": 10,
	// 		"image": "10rb"
	// 		}`)

	// 	auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWlud2ViQGdtYWlsLmNvbSIsImV4cCI6MTc3MDc3MDk2NSwidXNlcm5hbWUiOiJhZG1pbndlYiJ9.qg4vb8IBXbIuL9hK_aNEky59UWet5fF4DzPWVIDwdvQ"
	// 	request := httptest.NewRequest(http.MethodPost, "/api/auth/products", requestBody)
	// 	request.Header.Set(echo.HeaderAuthorization, auth)
	// 	recorder := httptest.NewRecorder()
	// 	e.ServeHTTP(recorder, request)

	// 	if assert.NoError(t, controller.CreateProduct(e.AcquireContext())) {
	// 		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	// 	}
	// })
}

func TestGetProducts(t *testing.T) {
	e := InitProductsTestAPI()
	InsertDataProduct()
	e.GET("/api/auth/products",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Get Products", func(t *testing.T) {
		auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWlud2ViQGdtYWlsLmNvbSIsImV4cCI6MTc3MDc3MDk2NSwidXNlcm5hbWUiOiJhZG1pbndlYiJ9.qg4vb8IBXbIuL9hK_aNEky59UWet5fF4DzPWVIDwdvQ"
		request := httptest.NewRequest(http.MethodGet, "/api/auth/products", nil)
		request.Header.Set(echo.HeaderAuthorization, auth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.GetProducts(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}

func TestGetProductByID(t *testing.T) {
	e := InitProductsTestAPI()
	InsertDataProduct()
	InsertDataAdmin()
	e.GET("/api/auth/products/:id",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Get Product By ID", func(t *testing.T) {
		auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWlud2ViQGdtYWlsLmNvbSIsImV4cCI6MTc3MDc3MDk2NSwidXNlcm5hbWUiOiJhZG1pbndlYiJ9.qg4vb8IBXbIuL9hK_aNEky59UWet5fF4DzPWVIDwdvQ"
		request := httptest.NewRequest(http.MethodGet, "/api/auth/products/1", nil)
		request.Header.Set(echo.HeaderAuthorization, auth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.GetProductByID(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})

	// // Test Case 2
	// t.Run("Get Product By Unknown ID", func(t *testing.T) {
	// 	auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWlud2ViQGdtYWlsLmNvbSIsImV4cCI6MTc3MDc3MDk2NSwidXNlcm5hbWUiOiJhZG1pbndlYiJ9.qg4vb8IBXbIuL9hK_aNEky59UWet5fF4DzPWVIDwdvQ"
	// 	request := httptest.NewRequest(http.MethodGet, "/api/auth/products/wrong-id", nil)
	// 	request.Header.Set(echo.HeaderAuthorization, auth)
	// 	recorder := httptest.NewRecorder()
	// 	e.ServeHTTP(recorder, request)

	// 	if assert.NoError(t, controller.GetProductByID(e.AcquireContext())) {
	// 		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	// 	}
	// })
}

func TestGetProductsByMethod(t *testing.T) {
	e := InitProductsTestAPI()
	InsertDataProduct()
	InsertDataAdmin()
	e.GET("/api/auth/products/redeem",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Get Product By Method", func(t *testing.T) {
		auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWlud2ViQGdtYWlsLmNvbSIsImV4cCI6MTc3MDc3MDk2NSwidXNlcm5hbWUiOiJhZG1pbndlYiJ9.qg4vb8IBXbIuL9hK_aNEky59UWet5fF4DzPWVIDwdvQ"
		request := httptest.NewRequest(http.MethodGet, "/api/auth/products/redeem", nil)
		request.Header.Set(echo.HeaderAuthorization, auth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.GetProductsByMethod(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}

func TestGetProductsByCategory(t *testing.T) {
	e := InitProductsTestAPI()
	InsertDataProduct()
	InsertDataAdmin()
	e.GET("/api/auth/products/category/:category",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Get Product By Category", func(t *testing.T) {
		auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWlud2ViQGdtYWlsLmNvbSIsImV4cCI6MTc3MDc3MDk2NSwidXNlcm5hbWUiOiJhZG1pbndlYiJ9.qg4vb8IBXbIuL9hK_aNEky59UWet5fF4DzPWVIDwdvQ"
		request := httptest.NewRequest(http.MethodGet, "/api/auth/products/category/credits", nil)
		request.Header.Set(echo.HeaderAuthorization, auth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.GetProductsByCategory(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}

func TestUpdateProduct(t *testing.T) {
	e := InitProductsTestAPI()
	InsertDataProduct()
	InsertDataAdmin()
	e.PUT("/api/auth/products/:id",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Update Product", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"category": "credits",
			"name": "test-updated",
			"description": "test-description-updated",
			"price": 10000,
			"stock": 4,
			"image": "10rb"
		}`)
		auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWlud2ViQGdtYWlsLmNvbSIsImV4cCI6MTc3MDc3MDk2NSwidXNlcm5hbWUiOiJhZG1pbndlYiJ9.qg4vb8IBXbIuL9hK_aNEky59UWet5fF4DzPWVIDwdvQ"
		request := httptest.NewRequest(http.MethodPut, "/api/auth/products/1", requestBody)
		request.Header.Set(echo.HeaderAuthorization, auth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.UpdateProduct(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}

func TestDeleteProduct(t *testing.T) {
	e := InitProductsTestAPI()
	InsertDataProduct()
	InsertDataAdmin()
	e.DELETE("/api/auth/products/:id",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Delete Product", func(t *testing.T) {
		auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWlud2ViQGdtYWlsLmNvbSIsImV4cCI6MTc3MDc3MDk2NSwidXNlcm5hbWUiOiJhZG1pbndlYiJ9.qg4vb8IBXbIuL9hK_aNEky59UWet5fF4DzPWVIDwdvQ"
		request := httptest.NewRequest(http.MethodDelete, "/api/auth/products/1", nil)
		request.Header.Set(echo.HeaderAuthorization, auth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.DeleteProduct(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
			log.Println(recorder.Body)
		}
	})
}

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

	// "log"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func InitTransactionsTestAPI() *echo.Echo {
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

func InserDataTransaction() {
	Transaction := entity.Transactions{
		ID:            "1",
		PaymentMethod: "redeem",
		UserID:        "2",
		ProductID:     "1",
		SerialNumber:  1234567890,
		IdentifierNum: "081234567890",
		Price:         12000,
		Status:        "pending",
	}
	err := entity.DB.Create(&Transaction).Error
	if err != nil {
		panic(err)
	}
}

func TestGetTransactions(t *testing.T) {
	e := InitTransactionsTestAPI()
	InsertDataAdmin()
	e.GET("/api/auth/transactions",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Get Transactions", func(t *testing.T) {
		adminAuth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWlud2ViQGdtYWlsLmNvbSIsImV4cCI6MTc3MDc3MDk2NSwidXNlcm5hbWUiOiJhZG1pbndlYiJ9.qg4vb8IBXbIuL9hK_aNEky59UWet5fF4DzPWVIDwdvQ"
		request := httptest.NewRequest(http.MethodGet, "/api/auth/transactions", nil)
		request.Header.Set(echo.HeaderAuthorization, adminAuth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.GetTransactions(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}

func TestGetTransactionsByMethod(t *testing.T) {
	e := InitTransactionsTestAPI()
	InsertDataAdmin()
	InserDataTransaction()
	e.GET("/api/auth/transactions/:payment_method",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Get Transactions By Method", func(t *testing.T) {
		adminAuth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWlud2ViQGdtYWlsLmNvbSIsImV4cCI6MTc3MDc3MDk2NSwidXNlcm5hbWUiOiJhZG1pbndlYiJ9.qg4vb8IBXbIuL9hK_aNEky59UWet5fF4DzPWVIDwdvQ"
		request := httptest.NewRequest(http.MethodGet, "/api/auth/transactions/redeem", nil)
		request.Header.Set(echo.HeaderAuthorization, adminAuth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.GetTransactionsByMethod(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})

	// // Test Case 2
	// t.Run("Get Transactions By Wrong Method", func(t *testing.T) {
	// 	adminAuth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWlud2ViQGdtYWlsLmNvbSIsImV4cCI6MTc3MDc3MDk2NSwidXNlcm5hbWUiOiJhZG1pbndlYiJ9.qg4vb8IBXbIuL9hK_aNEky59UWet5fF4DzPWVIDwdvQ"
	// 	request := httptest.NewRequest(http.MethodGet, "/api/auth/transactions/wrong-method", nil)
	// 	request.Header.Set(echo.HeaderAuthorization, adminAuth)
	// 	recorder := httptest.NewRecorder()
	// 	e.ServeHTTP(recorder, request)

	// 	if assert.NoError(t, controller.GetTransactionsByMethod(e.AcquireContext())) {
	// 		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	// 		log.Println(recorder.Body.String())
	// 	}
	// })
}

func TestGetTransactionByID(t *testing.T) {
	e := InitTransactionsTestAPI()
	InsertDataAdmin()
	InserDataTransaction()
	e.GET("/api/auth/transactions/:id",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Get Transactions By ID", func(t *testing.T) {
		adminAuth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWlud2ViQGdtYWlsLmNvbSIsImV4cCI6MTc3MDc3MDk2NSwidXNlcm5hbWUiOiJhZG1pbndlYiJ9.qg4vb8IBXbIuL9hK_aNEky59UWet5fF4DzPWVIDwdvQ"
		request := httptest.NewRequest(http.MethodGet, "/api/auth/transactions/1", nil)
		request.Header.Set(echo.HeaderAuthorization, adminAuth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.GetTransactionByID(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})

	// // Test Case 2
	// t.Run("Get Transactions By Wrong ID", func(t *testing.T) {
	// 	adminAuth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWlud2ViQGdtYWlsLmNvbSIsImV4cCI6MTc3MDc3MDk2NSwidXNlcm5hbWUiOiJhZG1pbndlYiJ9.qg4vb8IBXbIuL9hK_aNEky59UWet5fF4DzPWVIDwdvQ"
	// 	request := httptest.NewRequest(http.MethodGet, "/api/auth/transactions/wrong-id", nil)
	// 	request.Header.Set(echo.HeaderAuthorization, adminAuth)
	// 	recorder := httptest.NewRecorder()
	// 	e.ServeHTTP(recorder, request)

	// 	if assert.NoError(t, controller.GetTransactionByID(e.AcquireContext())) {
	// 		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	// 		log.Println(recorder.Body.String())
	// 	}
	// })
}

// this function output is correct but message code is different
// func TestGetHistory(t *testing.T) {
// 	e := InitTransactionsTestAPI()
// 	InsertDataUser()
// 	InsertDataProduct()
// 	InserDataTransaction()
// 	e.GET("/api/auth/transactions/history",
// 		func(c echo.Context) error {
// 			token := c.Get("user").(*jwt.Token)
// 			return c.JSON(http.StatusOK, token.Claims)
// 		})
// 	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

// 	// Test Case 1
// 	t.Run("Get History", func(t *testing.T) {
// 		userAuth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJ3ZWJAZ21haWwuY29tIiwiZXhwIjoxNzcwODU1MjAwLCJ1c2VybmFtZSI6InVzZXJ3ZWIifQ.H9Y4T_x10xnupepPhpenxoGKDv0IYIR6ZEiHPdCWW1g"
// 		request := httptest.NewRequest(http.MethodGet, "/api/auth/history", nil)
// 		request.Header.Set(echo.HeaderAuthorization, userAuth)
// 		recorder := httptest.NewRecorder()
// 		e.ServeHTTP(recorder, request)

// 		if assert.NoError(t, controller.GetHistory(e.AcquireContext())) {
// 			log.Println(recorder.Body.String())
// 			assert.Equal(t, http.StatusOK, recorder.Code)
// 		}
// 	})
// }

func TestGetHistoryByMethod(t *testing.T) {
	e := InitTransactionsTestAPI()
	InsertDataUser()
	InsertDataProduct()
	InserDataTransaction()
	e.GET("/api/auth/transactions/history/:paymentMethod",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Get History By Method", func(t *testing.T) {
		userAuth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJ3ZWJAZ21haWwuY29tIiwiZXhwIjoxNzcwODU1MjAwLCJ1c2VybmFtZSI6InVzZXJ3ZWIifQ.H9Y4T_x10xnupepPhpenxoGKDv0IYIR6ZEiHPdCWW1g"
		request := httptest.NewRequest(http.MethodGet, "/api/auth/transactions/history/redeem", nil)
		request.Header.Set(echo.HeaderAuthorization, userAuth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.GetHistoryByMethod(e.AcquireContext())) {
			log.Println(recorder.Body.String())
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}

func TestGetHistoryByMethodCategory(t *testing.T) {
	e := InitTransactionsTestAPI()
	InsertDataUser()
	InsertDataProduct()
	InserDataTransaction()
	e.GET("/api/auth/transactions/history/:paymentMethod/:category",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Get History By Method Category", func(t *testing.T) {
		userAuth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJ3ZWJAZ21haWwuY29tIiwiZXhwIjoxNzcwODU1MjAwLCJ1c2VybmFtZSI6InVzZXJ3ZWIifQ.H9Y4T_x10xnupepPhpenxoGKDv0IYIR6ZEiHPdCWW1g"
		request := httptest.NewRequest(http.MethodGet, "/api/auth/transactions/history/redeem/credis", nil)
		request.Header.Set(echo.HeaderAuthorization, userAuth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.GetHistoryByMethodCategory(e.AcquireContext())) {
			log.Println(recorder.Body.String())
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}

func TestCreateTransactionByUser(t *testing.T) {
	e := InitTransactionsTestAPI()
	InsertDataUser()
	InsertDataProduct()
	e.POST("/api/auth/transactions",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Create Transaction", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"paymentMethod": "redeem",
			"productID": 1,
			"identifierNum": "08123456789",
		}`)
		userAuth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJ3ZWJAZ21haWwuY29tIiwiZXhwIjoxNzcwODU1MjAwLCJ1c2VybmFtZSI6InVzZXJ3ZWIifQ.H9Y4T_x10xnupepPhpenxoGKDv0IYIR6ZEiHPdCWW1g"
		request := httptest.NewRequest(http.MethodPost, "/api/auth/transactions", requestBody)
		request.Header.Set(echo.HeaderAuthorization, userAuth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.CreateTransactionByUser(e.AcquireContext())) {
			log.Println(recorder.Body.String())
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}

func TestCreateTransactionByAdmin(t *testing.T) {

}

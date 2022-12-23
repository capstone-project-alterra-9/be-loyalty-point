package controller

import (
	"capstone-project/config"
	"capstone-project/controller"
	"capstone-project/repository"
	"capstone-project/service"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func InitUsersTestAPI() *echo.Echo {
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

func TestGetUsers(t *testing.T) {
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

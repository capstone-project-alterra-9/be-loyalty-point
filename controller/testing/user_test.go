package controller

import (
	"capstone-project/config"
	"capstone-project/controller"
	"capstone-project/entity"
	"capstone-project/repository"
	"capstone-project/service"
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

func InitUsersTestAPI() *echo.Echo {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
	db := config.InitDatabase()
	repository := repository.NewRepository(db)
	Service := service.NewService(repository)
	controller.NewController(Service)
	e := echo.New()
	return e
}

func TestGetUsers(t *testing.T) {
	e := InitUsersTestAPI()
	e.GET("/api/users",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Get Users", func(t *testing.T) {
		auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFnaS50ZXN0QGdtYWlsLmNvbSIsImV4cCI6MTY3MTc5ODQ3NCwidXNlcm5hbWUiOiJhZ2lwcm9kdWN0aW9uIn0.BSl2OssGIE0PAlf3z4Z8RDTeEzaVzJvlVcDpUZij_DQ"
		request := httptest.NewRequest(http.MethodGet, "/api/users", nil)
		request.Header.Set(echo.HeaderAuthorization, auth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.GetUsersPagination(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}

func TestCreateUserByAdmin(t *testing.T) {
	e := InitUsersTestAPI()
	e.POST("/api/users/create",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Create User by admin", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"username" : "test12345",
			"email" : "test356@gmail.com",
			"password" : "testpassword01",
			"point": 1000
		}`)

		auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFnaS50ZXN0QGdtYWlsLmNvbSIsImV4cCI6MTY3MTc5ODQ3NCwidXNlcm5hbWUiOiJhZ2lwcm9kdWN0aW9uIn0.BSl2OssGIE0PAlf3z4Z8RDTeEzaVzJvlVcDpUZij_DQ"
		request := httptest.NewRequest(http.MethodPost, "/api/users/create", requestBody)
		request.Header.Set(echo.HeaderAuthorization, auth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.CreateUserByAdmin(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}

func TestGetUserByID(t *testing.T) {
	e := InitUsersTestAPI()
	e.GET("/api/users/:id",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Get User By ID", func(t *testing.T) {
		auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFnaS50ZXN0QGdtYWlsLmNvbSIsImV4cCI6MTY3MTc5ODQ3NCwidXNlcm5hbWUiOiJhZ2lwcm9kdWN0aW9uIn0.BSl2OssGIE0PAlf3z4Z8RDTeEzaVzJvlVcDpUZij_DQ"
		request := httptest.NewRequest(http.MethodGet, "/api/users/aaecd930-e491-4010-965e-f715abdf7144", nil)
		request.Header.Set(echo.HeaderAuthorization, auth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.GetOneByUserId(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}

func TestUpdateUserById(t *testing.T) {
	e := InitUsersTestAPI()
	e.PUT("/api/users/:id",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Update Product", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"role" : "user",
			"username" : "saka",
			"email" : "testmail01@gmail.com",
			"password" : "testpassword01",
			"points" : 30000,
			"costPoints" : 0
		}`)
		auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFnaS50ZXN0QGdtYWlsLmNvbSIsImV4cCI6MTY3MTc5ODQ3NCwidXNlcm5hbWUiOiJhZ2lwcm9kdWN0aW9uIn0.BSl2OssGIE0PAlf3z4Z8RDTeEzaVzJvlVcDpUZij_DQ"
		request := httptest.NewRequest(http.MethodPut, "/api/users/aaecd930-e491-4010-965e-f715abdf7144", requestBody)
		request.Header.Set(echo.HeaderAuthorization, auth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.UpdateOneByUserId(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}

func TestDeleteUserById(t *testing.T) {
	e := InitUsersTestAPI()
	e.DELETE("/api/users/:id",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Delete User By Id", func(t *testing.T) {
		auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFnaS50ZXN0QGdtYWlsLmNvbSIsImV4cCI6MTY3MTc5ODQ3NCwidXNlcm5hbWUiOiJhZ2lwcm9kdWN0aW9uIn0.BSl2OssGIE0PAlf3z4Z8RDTeEzaVzJvlVcDpUZij_DQ"
		request := httptest.NewRequest(http.MethodDelete, "/api/users/aaecd930-e491-4010-965e-f715abdf7144", nil)
		request.Header.Set(echo.HeaderAuthorization, auth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.DeleteOneById(e.AcquireContext())) {
			// log.Println(recorder.Body)
			assert.Equal(t, http.StatusOK, recorder.Code)
			var users []entity.Users
			err := entity.DB.Find(&users, "role = ?", "user").Error
			assert.NoError(t, err)
			assert.Equal(t, 0, len(users))
		}
	})
}

func TestCountGetUsers(t *testing.T) {
	e := InitUsersTestAPI()
	e.GET("/api/users/count",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Get Count Users", func(t *testing.T) {
		auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFnaS50ZXN0QGdtYWlsLmNvbSIsImV4cCI6MTY3MTc5ODQ3NCwidXNlcm5hbWUiOiJhZ2lwcm9kdWN0aW9uIn0.BSl2OssGIE0PAlf3z4Z8RDTeEzaVzJvlVcDpUZij_DQ"
		request := httptest.NewRequest(http.MethodGet, "/api/users/count", nil)
		request.Header.Set(echo.HeaderAuthorization, auth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.GetCountUsers(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}

func TestSendEmailForgotPassword(t *testing.T) {
	e := InitUsersTestAPI()
	e.POST("/api/forgot-password",
		func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			return c.JSON(http.StatusOK, token.Claims)
		})
	e.Use(mid.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// Test Case 1
	t.Run("Create User by admin", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"email": "ajizapar080500@gmail.com"
		}`)

		auth := "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFnaS50ZXN0QGdtYWlsLmNvbSIsImV4cCI6MTY3MTc5ODQ3NCwidXNlcm5hbWUiOiJhZ2lwcm9kdWN0aW9uIn0.BSl2OssGIE0PAlf3z4Z8RDTeEzaVzJvlVcDpUZij_DQ"
		request := httptest.NewRequest(http.MethodPost, "/api/forgot-password", requestBody)
		request.Header.Set(echo.HeaderAuthorization, auth)
		recorder := httptest.NewRecorder()
		e.ServeHTTP(recorder, request)

		if assert.NoError(t, controller.SendEmailForgotPassword(e.AcquireContext())) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}

package controller

import (
	"capstone-project/config"
	"capstone-project/controller"
	"capstone-project/entity"
	"capstone-project/repository"
	"capstone-project/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func InitAuthTestAPI() *echo.Echo {
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

func InsertDataAdmin() {
	password, _ := bcrypt.GenerateFromPassword([]byte("adminweb567"), bcrypt.DefaultCost)
	Admin := entity.Users{
		ID:       "1",
		Role:     "admin",
		Username: "adminweb",
		Email:    "adminweb@gmail.com",
		Password: string(password),
	}
	err := entity.DB.Create(&Admin).Error
	if err != nil {
		panic(err)
	}
}

func InsertDataUser() {
	password, _ := bcrypt.GenerateFromPassword([]byte("userweb567"), bcrypt.DefaultCost)
	User := entity.Users{
		ID:       "2",
		Role:     "user",
		Username: "userweb",
		Email:    "userweb@gmail.com",
		Password: string(password),
	}
	err := entity.DB.Create(&User).Error
	if err != nil {
		panic(err)
	}
}

func TestLogin(t *testing.T) {
	e := InitAuthTestAPI()
	InsertDataAdmin()
	InsertDataUser()
	// Test Case 1
	t.Run("Login Admin", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"email": "adminweb@gmail.com",
			"password": "adminweb567"
			}`)
		request := httptest.NewRequest(http.MethodPost, "/api/login", requestBody)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		c := e.NewContext(request, recorder)

		if assert.NoError(t, controller.Login(c)) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})

	// Test Case 2
	t.Run("Login User", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"email": "wrong-user@gmail.com",
			"password": "wrong-password"
		}`)
		request := httptest.NewRequest(http.MethodPost, "/api/login", requestBody)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		c := e.NewContext(request, recorder)

		if assert.NoError(t, controller.Login(c)) {
			assert.Equal(t, http.StatusUnauthorized, recorder.Code)
		}
	})
}

func TestRegister(t *testing.T) {
	e := InitAuthTestAPI()
	// Test Case 1
	t.Run("Register 1", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"username": "adminweb01",
			"email": "adminweb@gmail.com",
			"password": "adminweb567"
		}`)
		request := httptest.NewRequest(http.MethodPost, "/api/register", requestBody)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		c := e.NewContext(request, recorder)

		if assert.NoError(t, controller.Register(c)) {
			assert.Equal(t, http.StatusCreated, recorder.Code)
		}
	})

	// Test Case 2
	t.Run("Register 2", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"username": "kiw",
			"email": "bad-email",
			"password": "kiw123"
		}`)
		request := httptest.NewRequest(http.MethodPost, "/api/register", requestBody)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		c := e.NewContext(request, recorder)

		if assert.NoError(t, controller.Register(c)) {
			assert.Equal(t, http.StatusBadRequest, recorder.Code)
		}
	})
}

func TestGenerateRefreshToken(t *testing.T) {
	e := InitAuthTestAPI()
	// Test Case 1
	t.Run("Generate Refresh Token", func(t *testing.T) {
		requestBody := strings.NewReader(`{
			"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWlud2ViQGdtYWlsLmNvbSIsImV4cCI6MTY3MDc3MDk2NSwidXNlcm5hbWUiOiJhZG1pbndlYiJ9.vvVtcnUObaAQwgkCIo2i-jQtoGLWog41hhRuYxefN8o"
		}`)
		request := httptest.NewRequest(http.MethodPost, "/api/refresh-token", requestBody)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		c := e.NewContext(request, recorder)

		if assert.NoError(t, controller.GenerateRefreshToken(c)) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}

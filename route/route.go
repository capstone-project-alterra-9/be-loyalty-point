package route

import (
	"capstone-project/controller"
	m "capstone-project/middleware"
	"capstone-project/service"
	"os"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New(Service service.Svc) *echo.Echo {
	controller.NewController(Service)
	e := echo.New()
	e.Use(mid.CORSWithConfig(mid.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	m.LogMiddleware(e)

	// Routing withouth JWT
	// e.POST("/login", controller.Login)
	// e.POST("/register", controller.Register)

	e.Use(mid.JWT([]byte(os.Getenv("SECRET_JWT"))))
	// Routing with JWT

	return e
}
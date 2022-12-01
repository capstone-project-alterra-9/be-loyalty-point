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
	eApi := e.Group("/api")
	eApi.POST("/login", controller.Login)
	eApi.POST("/register", controller.Register)

	eAuth := eApi.Group("/auth")
	eAuth.Use(mid.JWT([]byte(os.Getenv("SECRET_JWT"))))
	// Routing with JWT
	eAuth.GET("/transactions", controller.GetTransactions)
	eAuth.GET("/transactions/redeem", controller.GetTransactionsRedeem)
	eAuth.GET("/transactions/buy", controller.GetTransactionsBuy)
	eAuth.GET("/transactions/:id", controller.GetTransactionByID)
	eAuth.GET("/history", controller.GetTransactionsByUser)
	eAuth.POST("/transactions", controller.CreateTransactionByUser)
	eAuth.POST("/transactions/dummy", controller.CreateTransactionByAdmin)
	eAuth.PUT("/transactions/:id", controller.UpdateTransactionByAdmin)
	eAuth.DELETE("/transactions/:id", controller.DeleteTransactionByAdmin)

	// User endpoint
	eAuth.DELETE("/users", controller.DeleteTransactionByAdmin)

	return e
}

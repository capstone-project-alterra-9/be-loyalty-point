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
	eAuth.GET("/transactions/method/:paymentMethod", controller.GetTransactionsByMethod)
	eAuth.GET("/transactions/:id", controller.GetTransactionByID)
	eAuth.GET("/history", controller.GetHistory)
	eAuth.GET("/history/method/:paymentMethod", controller.GetHistoryByMethod)
	eAuth.GET("/history/:id", controller.GetTransactionByID)
	eAuth.POST("/transactions", controller.CreateTransactionByUser)
	eAuth.POST("/transactions/dummy", controller.CreateTransactionByAdmin)
	eAuth.PUT("/transactions/:id", controller.UpdateTransactionByAdmin)
	eAuth.DELETE("/transactions/:id", controller.DeleteTransactionByAdmin)

	eAuth.POST("/products", controller.CreateProduct)
	eAuth.GET("/products", controller.GetProducts)
	eAuth.GET("/products/category/:categoryName", controller.GetProductsByCategory)
	eAuth.GET("/products/:id", controller.GetProductByID)
	eAuth.PUT("/products/:id", controller.UpdateProduct)
	eAuth.DELETE("/products/:id", controller.DeleteProduct)

	// User endpoint
	eUser := eApi.Group("/users")
	eUser.Use(mid.JWT([]byte(os.Getenv("SECRET_JWT"))))
	eUser.DELETE("/:id", controller.DeleteOneById)
	eUser.GET("/:id", controller.GetOneByUserId)
	eUser.GET("/users", controller.GetUsersPagination)

	return e
}

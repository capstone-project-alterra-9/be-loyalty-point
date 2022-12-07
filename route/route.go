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
	eAuth.GET("/history", controller.GetHistory)
	eAuth.GET("/history/method/:paymentMethod", controller.GetHistoryByMethod)
	eAuth.GET("/history/:id", controller.GetTransactionByID)

	eTransaction := eAuth.Group("/transactions")
	eTransaction.GET("", controller.GetTransactions)
	eTransaction.GET("/method/:paymentMethod", controller.GetTransactionsByMethod)
	eTransaction.GET("/:id", controller.GetTransactionByID)
	eTransaction.POST("", controller.CreateTransactionByUser)
	eTransaction.POST("/dummy", controller.CreateTransactionByAdmin)
	eTransaction.PUT("/:id", controller.UpdateTransactionByAdmin)
	eTransaction.DELETE("/:id", controller.DeleteTransactionByAdmin)

	eProduct := eAuth.Group("/products")
	eProduct.POST("", controller.CreateProduct)
	eProduct.GET("", controller.GetProducts)
	eProduct.GET("/category/:categoryName", controller.GetProductsByCategory)
	eProduct.GET("/:id", controller.GetProductByID)
	eProduct.PUT("/:id", controller.UpdateProduct)
	eProduct.DELETE("/:id", controller.DeleteProduct)

	// User endpoint
	eUser := eApi.Group("/users")
	eUser.Use(mid.JWT([]byte(os.Getenv("SECRET_JWT"))))
	eUser.DELETE("/:id", controller.DeleteOneById)
	eUser.GET("/:id", controller.GetOneByUserId)
	eUser.GET("/users", controller.GetUsersPagination)
	eUser.PUT("/:id", controller.UpdateOneByUserId)
	eUser.GET("", controller.GetUsersPagination)
	eUser.POST("/create", controller.CreateUserByAdmin)
	// eUser.GET("/count", controller.GetCountUsers)

	return e
}

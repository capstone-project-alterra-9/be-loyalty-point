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
		Skipper:      mid.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	m.LogMiddleware(e)

	e.GET("/", controller.Connected)
	eApi := e.Group("/api")
	eApi.POST("/login", controller.Login)
	eApi.POST("/register", controller.Register)
	eApi.POST("/refresh-token", controller.GenerateRefreshToken)

	eAuth := eApi.Group("/auth")
	eAuth.Use(mid.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eHistory := eAuth.Group("/history")
	eHistory.GET("", controller.GetHistory)
	eHistory.GET("/method/:paymentMethod", controller.GetHistoryByMethod)
	eHistory.GET("/method/:paymentMethod/:category", controller.GetHistoryByMethodCategory)
	eHistory.GET("/:id", controller.GetTransactionByID)

	eTransaction := eAuth.Group("/transactions")
	eTransaction.GET("", controller.GetTransactions)
	eTransaction.GET("/method/:paymentMethod", controller.GetTransactionsByMethod)
	eTransaction.GET("/:id", controller.GetTransactionByID)
	eTransaction.POST("", controller.CreateTransactionByUser)
	eTransaction.POST("/admin/create", controller.CreateTransactionByAdmin)
	eTransaction.PUT("/:id", controller.UpdateTransactionByAdmin)
	eTransaction.DELETE("/:id", controller.DeleteTransactionByAdmin)
	eTransaction.GET("/count", controller.GetCountTransactions)
	eTransaction.POST("/midtrans", controller.CreateMidtransTransaction)
	eTransaction.POST("/payment-status/:status", controller.HandlingPaymentStatus)

	eProduct := eAuth.Group("/products")
	eProduct.POST("", controller.CreateProduct)
	eProduct.GET("", controller.GetProducts)
	eProduct.GET("/method/:paymentMethod", controller.GetProductsByMethod)
	eProduct.GET("/category/:categoryName", controller.GetProductsByCategory)
	eProduct.GET("/:id", controller.GetProductByID)
	eProduct.PUT("/:id", controller.UpdateProduct)
	eProduct.DELETE("/:id", controller.DeleteProduct)
	eProduct.GET("/count", controller.GetCountProducts)

	eUser := eApi.Group("/users")
	eUser.Use(mid.JWT([]byte(os.Getenv("SECRET_JWT"))))
	eUser.DELETE("/:id", controller.DeleteOneById)
	eUser.GET("/:id", controller.GetOneByUserId)
	eUser.PUT("/:id", controller.UpdateOneByUserId)
	eUser.GET("", controller.GetUsersPagination)
	eUser.POST("/create", controller.CreateUserByAdmin)
	eUser.GET("/count", controller.GetCountUsers)
	eApi.POST("/forgot-password", controller.SendEmailForgotPassword)

	eFaq := eApi.Group("/faqs")
	eFaq.Use(mid.JWT([]byte(os.Getenv("SECRET_JWT"))))
	eFaq.GET("", controller.GetFaqs)
	eFaq.POST("/create", controller.CreateFaq)

	return e
}

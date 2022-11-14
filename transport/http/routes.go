package httpRoutes

import (
	"net/http"

	"github.com/aviralbansal29/bill_split/app/api/controllers"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// SetupRouter defines url endpoints
func SetupRouter(router *echo.Echo) {
	router.Pre(echoMiddleware.RemoveTrailingSlash())
	router.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Status OK"})
	})

	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.RetrieveUser)

	router.POST("/groups", controllers.CreateGroup)
	router.POST("/groups/:group_id/add-user/:user_id", controllers.AddUser)

	router.GET("/transactions", controllers.GetTransaction)
	router.POST("/transactions", controllers.CreateTransaction)

	router.GET("/swagger/*any", echoSwagger.WrapHandler)
}

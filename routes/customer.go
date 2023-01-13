package routes

import (
	"github.com/labstack/echo/v4"
	"sample-api-rest/controllers"
)

func Customer(e *echo.Echo) {
	r := e.Group("/customers")
	/*middlewares*/

	/*routes*/
	r.GET("", controllers.CustomerController.FindAll)
	r.GET("/:id", controllers.CustomerController.FindById)
}

package routes

import (
	"github.com/Rayato159/isekai-shop-api/app/controllers"
	"github.com/labstack/echo/v4"
)

func NewRoutes(e *echo.Echo) {
	postControllers := controllers.NewPostControllers()

	e.GET("/posts/:post_id", postControllers.FindOnePost)
}

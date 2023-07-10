package routes

import (
	"landtick/handlers"
	"landtick/pkg/middleware"
	"landtick/pkg/mysql"
	"landtick/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)
	e.GET("/users", h.FindUser)
	e.GET("/channel/:id", h.GetUser)
	e.POST("/user", middleware.Auth(h.CreateUser))
	e.PATCH("/user/:id", middleware.Auth(h.UpdateUser))
	e.DELETE("/user/:id", middleware.Auth(h.DeleteUser))
}

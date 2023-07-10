package routes

import (
	"landtick/handlers"
	"landtick/pkg/mysql"
	"landtick/repositories"

	"github.com/labstack/echo/v4"
)

func TicketRoutes(e *echo.Group) {
	ticketRepository := repositories.RepositoryTicket(mysql.DB)
	h := handlers.HandlerTicket(ticketRepository)
	e.GET("/tickets", h.FindTicket)
	e.GET("/tickets/:id", h.GetTicket)
}

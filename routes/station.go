package routes

import (
	"landtick/handlers"
	"landtick/pkg/mysql"
	"landtick/repositories"

	"github.com/labstack/echo/v4"
)

func StationRoutes(e *echo.Group) {
	stationRepository := repositories.RepositoryStation(mysql.DB)
	h := handlers.HandlerStation(stationRepository)

	e.GET("/stations", h.FindStation) // Menampilkan Semua data Station
	e.GET("/station/:id", h.GetStation)
	e.POST("/stations", h.CreateStation)
	e.PATCH("/stations/:id", h.UpdateStation)
	e.DELETE("/stations/:id", h.DeleteStation)
}

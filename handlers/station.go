package handlers

import (
	dto "landtick/dto/result"
	stationsdto "landtick/dto/stations"
	"landtick/models"
	"landtick/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerStation struct {
	StationRepository repositories.StationRepository
}

type dataStation struct {
	Station interface{} `json:"stations"`
}

func HandlerStation(StationRepository repositories.StationRepository) *handlerStation {
	return &handlerStation{StationRepository}
}

// Menampilkan Semua data Station
func (h *handlerStation) FindStation(c echo.Context) error {
	stations, err := h.StationRepository.FindStation()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: stations})
}

// Menampilkan Semua data stasiun Berdasarkan ID
func (h *handlerStation) GetStation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	station, err := h.StationRepository.GetStation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: dataStation{Station: convertResponseStation(station)}})
}

func (h *handlerStation) CreateStation(c echo.Context) error {
	request := new(stationsdto.CreateStationRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	station := models.Station{
		Name: request.Name,
	}

	data, err := h.StationRepository.CreateStation(station)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertResponseStation(data)})

}

func (h *handlerStation) UpdateStation(c echo.Context) error {
	request := new(stationsdto.UpdateStationRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	id, _ := strconv.Atoi(c.Param("id"))

	station, err := h.StationRepository.GetStation(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	if request.Name != "" {
		station.Name = request.Name
	}

	data, err := h.StationRepository.UpdateStation(station)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertResponseStation(data)})
}

func (h *handlerStation) DeleteStation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	station, err := h.StationRepository.GetStation(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	data, err := h.StationRepository.DeleteStation(station)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertResponseStation(data)})
}
func convertResponseStation(u models.Station) stationsdto.StationResponse {
	return stationsdto.StationResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}

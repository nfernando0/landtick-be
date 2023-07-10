package handlers

import (
	resultdto "landtick/dto/result"
	ticketdto "landtick/dto/ticket"
	"landtick/models"
	"landtick/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handlerTicket struct {
	TicketRepository repositories.TicketRepository
}

func HandlerTicket(TicketRepository repositories.TicketRepository) *handlerTicket {
	return &handlerTicket{TicketRepository}
}

// Menampilkan Semua data tiket
func (h *handlerTicket) FindTicket(c echo.Context) error {
	tickets, err := h.TicketRepository.FindTicket()
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	response := make([]ticketdto.TicketResponseGet, len(tickets))
	for i, t := range tickets {
		response[i] = convertResponseTicket(t)
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: "Success", Data: response})
}

// Menampilkan data tiket berdasarkan ID
func (h *handlerTicket) GetTicket(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	ticket, err := h.TicketRepository.GetTicket(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: "Success", Data: convertResponseTicket(ticket)})
}

func convertResponseTicket(t models.Ticket) ticketdto.TicketResponseGet {
	return ticketdto.TicketResponseGet{
		ID:                   t.ID,
		NameTrain:            t.NameTrain,
		TypeTrain:            t.TypeTrain,
		StartDate:            t.StartTime.Format("2006-01-02 15:04:05"),
		StartStationID:       t.StartStation.ID,
		StartTime:            t.StartTime.Format("15:04"),
		ArrivalTime:          t.StartTime.Format("15:04"),
		DestinationStationID: t.DestinationStation.ID,
		Price:                t.Price,
		Qty:                  t.Qty,
	}
}

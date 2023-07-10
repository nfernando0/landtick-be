package ticketdto

type CreateTicketRequest struct {
	NameTrain            string `json:"name_train" form:"name_train" validate:"required"`
	TypeTrain            string `json:"type_train" form:"type_train" validate:"required"`
	StartDate            string `json:"start_date" form:"start_date" validate:"required"`
	StartTime            string `json:"start_time" form:"start_time" validate:"required"`
	StartStationID       int    `json:"start_station" form:"start_station" validate:"required"`
	ArrivalTime          string `json:"arrival_time" form:"arrival_time" validate:"required"`
	DestinationStationID int    `json:"destination_station" form:"destination_station" validate:"required"`
	Price                int    `json:"price" form:"price" validate:"required"`
	Qty                  int    `json:"qty" form:"qty" validate:"required"`
}

type UpdateTicketRequest struct {
	NameTrain      string `json:"name_train" form:"name_train" validate:"required"`
	TypeTrain      string `json:"type_train" form:"type_train" validate:"required"`
	StartDate      string `json:"start_date" form:"start_date" validate:"required"`
	StartTime      string `json:"start_time" form:"start_time" validate:"required"`
	StartStationID int    `json:"start_station" form:"start_station" validate:"required"`
	ArrivalTime    string `json:"arrival_time" form:"arrival_time" validate:"required"`
	Price          int    `json:"price" form:"price" validate:"required"`
	Qty            int    `json:"qty" form:"qty" validate:"required"`
}

type CreateTicketResponse struct {
	ID                   int    `json:"id"`
	NameTrain            string `json:"name_train" validate:"required"`
	TypeTrain            string `json:"type_train" validate:"required"`
	StartDate            string `json:"start_date" validate:"required"`
	StartStationID       int    `json:"start_station_id"`
	StartTime            string `json:"start_time" validate:"required"`
	ArrivalTime          string `json:"arrival_time" validate:"required"`
	DestinationStationID int    `json:"destination_station_id"`
	Price                int    `json:"price" validate:"required"`
	Qty                  int    `json:"qty" validate:"required"`
}

type TicketResponseGet struct {
	ID                   int    `json:"id"`
	NameTrain            string `json:"name_train" validate:"required"`
	TypeTrain            string `json:"type_train" validate:"required"`
	StartDate            string `json:"start_date" validate:"required"`
	StartStationID       int    `json:"start_station_id"`
	StartTime            string `json:"start_time" validate:"required"`
	ArrivalTime          string `json:"arrival_time" validate:"required"`
	DestinationStationID int    `json:"destination_station_id"`
	Price                int    `json:"price" validate:"required"`
	Qty                  int    `json:"qty" validate:"required"`
}

type MyTicketResponse struct {
	ID          int    `json:"id"`
	NameTrain   string `json:"name_train" validate:"required"`
	TypeTrain   string `json:"type_train" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
	StartTime   string `json:"start_time" validate:"required"`
	ArrivalTime string `json:"arrival_time" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Qty         int    `json:"qty" validate:"required"`
}

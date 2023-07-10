package models

import "time"

type Ticket struct {
	ID                   int       `json:"id"`
	NameTrain            string    `json:"name_train"`
	TypeTrain            string    `json:"type_train"`
	StartDate            time.Time `json:"start_date" gorm:"type:date"`
	StartStationID       int       `json:"start_station_id"`
	StartStation         Station   `json:"start_station"`
	StartTime            time.Time `json:"start_time" gorm:"type:time"`
	ArrivalTime          time.Time `json:"arrival_time" gorm:"type:time"`
	DestinationStationID int       `json:"destination_station_id"`
	DestinationStation   Station   `json:"destination_station"`
	Price                int       `json:"price"`
	Qty                  int       `json:"-"`
}

type TicketResponseModels struct {
	ID                     int                  `json:"id"`
	NameTrain              string               `json:"name_train"`
	TypeTrain              string               `json:"type_train"`
	StartDate              time.Time            `json:"start_date" gorm:"type:date"`
	Start_Station_id       int                  `json:"start_station_id"`
	StartStation           StationResponseModel `json:"start_station"`
	StartTime              time.Time            `json:"start_time" gorm:"type:time"`
	ArrivalTime            time.Time            `json:"arrival_time" gorm:"type:time"`
	Destination_Station_id int                  `json:"destination_station_id"`
	DestinationStation     StationResponseModel `json:"destination_station"`
	Price                  int                  `json:"price"`
	Qty                    int                  `json:"-"`
}

func (TicketResponseModels) TableName() string {
	return "tickets"
}

package stationsdto

type CreateStationRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type UpdateStationRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type StationResponse struct {
	ID    int    `json:"id"`
	Name string `json:"name" form:"name" validate:"required"`
}

package models

type Station struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"type:varchar(255)"`
}

type StationResponseModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type StationResponse struct {
	Name string `json:"name"`
}

func (StationResponseModel) TableName() string {
	return "stations"
}

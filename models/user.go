package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullname" gorm:"type:varchar(255)"`
	Username  string    `json:"username" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255)"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
type UserResponseTicketModels struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (UserResponse) TableName() string {
	return "users"
}

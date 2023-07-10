package handlers

import (
	dto "landtick/dto/result"
	usersdto "landtick/dto/users"
	"landtick/models"
	"landtick/repositories"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	UserRepository repositories.UserRepository
}

type dataUser struct {
	User interface{} `json:"users"`
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

// Menampilkan data Semua User
func (h *handler) FindUser(c echo.Context) error {
	users, err := h.UserRepository.FindUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: dataUser{User: users}})
}

// menampilkan data user berdasarkan Id
func (h *handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: dataUser{User: convertResponse(user)}})
}

// Membuat User Baru
func (h *handler) CreateUser(c echo.Context) error {
	request := new(usersdto.CreateUserRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	user := models.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	data, err := h.UserRepository.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertResponse(data)})
}

// Merubah data User yang sudah ada
func (h *handler) UpdateUser(c echo.Context) error {
	request := new(usersdto.UpdateUserRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	if request.Fullname != "" {
		user.Fullname = request.Fullname
	}
	if request.Username != "" {
		user.Username = request.Username
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Password != "" {
		user.Password = request.Password
	}

	data, err := h.UserRepository.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertResponse(data)})

}

// Menghapus data user berdasarkan ID
func (h *handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	data, err := h.UserRepository.DeleteUser(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertResponse(data)})
}

func convertResponse(u models.User) usersdto.UserResponseGet {
	return usersdto.UserResponseGet{
		Id:       u.ID,
		Fullname: u.Fullname,
		Username: u.Username,
		Email:    u.Email,
	}
}

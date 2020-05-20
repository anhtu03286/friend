package controller

import (
	"encoding/json"
	"github.com/anhtu03286/friend/dto"
	"github.com/anhtu03286/friend/response"
	"github.com/anhtu03286/friend/service"
	"net/http"
)

type UserController struct {
	UserService service.IUserService
}

func (u UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	friends := u.UserService.GetAllUsers()
	res := response.Response{Success: true, Friends: friends, Count: len(friends)}
	response.SuccessResponse(w, res)
}

func (u UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	emailDto := dto.EmailDto{}

	if err := json.NewDecoder(r.Body).Decode(&emailDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body!")
		return
	}

	result, err := u.UserService.CreateUser(emailDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}
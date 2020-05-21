package controller

import (
	"encoding/json"
	"net/http"

	"github.com/anhtu03286/friend/dto"
	"github.com/anhtu03286/friend/response"
	"github.com/anhtu03286/friend/service"
)

type FriendController struct {
	FriendService service.IFriendService
}

func (f FriendController) CreateFriend(w http.ResponseWriter, r *http.Request) {
	var friendDto dto.FriendDto
	if err := json.NewDecoder(r.Body).Decode(&friendDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body!")
		return
	}

	result, err := f.FriendService.CreateFriend(friendDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}

func (f FriendController) CreateSubscribe(w http.ResponseWriter, r *http.Request) {
	var requestDto dto.RequestDto
	if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body!")
		return
	}

	result, err := f.FriendService.CreateSubscribe(requestDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}

func (f FriendController) CreateBlock(w http.ResponseWriter, r *http.Request) {
	var requestDto dto.RequestDto
	if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body!")
		return
	}

	result, err := f.FriendService.CreateBlock(requestDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Success{Success: result}
	response.SuccessResponse(w, res)
}

func (f FriendController) GetFriendsListByEmail(w http.ResponseWriter, r *http.Request) {
	var emailDto dto.EmailDto
	if err := json.NewDecoder(r.Body).Decode(&emailDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body!")
		return
	}
	results, err := f.FriendService.GetFriendsListByEmail(emailDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Response{Success: true, Friends: results, Count: len(results)}
	response.SuccessResponse(w, res)
}

func (f FriendController) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	var friendDto dto.FriendDto
	if err := json.NewDecoder(r.Body).Decode(&friendDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body!")
		return
	}

	results, err := f.FriendService.GetCommonFriends(friendDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Response{Success: true, Friends: results, Count: len(results)}
	response.SuccessResponse(w, res)
}

func (f FriendController) GetReceiversList(w http.ResponseWriter, r *http.Request) {
	var senderDto dto.SenderDto
	if err := json.NewDecoder(r.Body).Decode(&senderDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body!")
		return
	}

	results, err := f.FriendService.GetReceiversList(senderDto)
	if err != nil {
		response.ErrorResponse(w, err.Code, err.Message)
		return
	}
	res := response.Response{Success: true, Friends: results, Count: len(results)}
	response.SuccessResponse(w, res)
}
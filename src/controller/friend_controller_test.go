package controller

import (
	"bytes"
	"encoding/json"
	"github.com/anhtu03286/friend/dto"
	"github.com/anhtu03286/friend/response"
	"github.com/anhtu03286/friend/service"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateFriend(t *testing.T)  {
	jsonStr := []byte(`{"friends":["a@gmail.com","b@gmail.com"]}`)

	var friendDto dto.FriendDto
	json.Unmarshal([]byte(jsonStr), &friendDto)

	friendServiceMock := service.FriendServiceMock{}
	friendServiceMock.On("CreateFriend", friendDto).Return(true)
	friendController := FriendController{friendServiceMock}

	r, _ := http.NewRequest("POST", "/friends/create-friend", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	handler := http.HandlerFunc(friendController.CreateFriend)
	handler.ServeHTTP(w, r)

	var actualResult response.Success
	body, _ := ioutil.ReadAll(w.Result().Body)
	json.Unmarshal(body, &actualResult)

	expectedResult := response.Success{Success: true}

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, expectedResult.Success, actualResult.Success)
}

func TestCreateSubscribe(t *testing.T)  {
	jsonStr := []byte(`{"requestor":"a@gmail.com","target":"b@gmail.com"}`)

	var requestDto dto.RequestDto
	json.Unmarshal([]byte(jsonStr), &requestDto)

	friendServiceMock := service.FriendServiceMock{}
	friendServiceMock.On("CreateSubscribe", requestDto).Return(true)
	friendController := FriendController{friendServiceMock}

	r, _ := http.NewRequest("POST", "/friends/subscribe", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	handler := http.HandlerFunc(friendController.CreateSubscribe)
	handler.ServeHTTP(w, r)

	var actualResult response.Success
	body, _ := ioutil.ReadAll(w.Result().Body)
	json.Unmarshal(body, &actualResult)

	expectedResult := response.Success{Success: true}

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, expectedResult.Success, actualResult.Success)
}

func TestCreateBlock(t *testing.T)  {
	jsonStr := []byte(`{"requestor":"a@gmail.com","target":"b@gmail.com"}`)

	var requestDto dto.RequestDto
	json.Unmarshal([]byte(jsonStr), &requestDto)

	friendServiceMock := service.FriendServiceMock{}
	friendServiceMock.On("CreateBlock", requestDto).Return(true)
	friendController := FriendController{friendServiceMock}

	r, _ := http.NewRequest("POST", "/friends/block", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	handler := http.HandlerFunc(friendController.CreateBlock)
	handler.ServeHTTP(w, r)

	var actualResult response.Success
	body, _ := ioutil.ReadAll(w.Result().Body)
	json.Unmarshal(body, &actualResult)

	expectedResult := response.Success{Success: true}

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, expectedResult.Success, actualResult.Success)
}

func TestGetFriendsListByEmail(t *testing.T)  {
	jsonStr := []byte(`{"email":"a@gmail.com'"}`)
	friends := []string{"b@gmail.com", "c@gmail.com"}

	var emailDto dto.EmailDto
	json.Unmarshal([]byte(jsonStr), &emailDto)

	friendServiceMock := service.FriendServiceMock{}
	friendServiceMock.On("GetFriendsListByEmail", emailDto).Return(friends)
	friendController := FriendController{friendServiceMock}

	r, _ := http.NewRequest("POST", "/friends/get-friends-list", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	handler := http.HandlerFunc(friendController.GetFriendsListByEmail)
	handler.ServeHTTP(w, r)

	var actualResult response.Response
	body, _ := ioutil.ReadAll(w.Result().Body)
	json.Unmarshal(body, &actualResult)

	expectedResult := response.Response{Success: true, Friends: friends, Count: len(friends)}

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, expectedResult.Success, actualResult.Success)
	assert.Equal(t, expectedResult.Friends, actualResult.Friends)
	assert.Equal(t, expectedResult.Count, actualResult.Count)
}

func TestGetCommonFriends(t *testing.T)  {
	jsonStr := []byte(`{"friends":["a@gmail.com","b@gmail.com"]}`)
	friends := []string{"c@gmail.com"}

	var friendDto dto.FriendDto
	json.Unmarshal([]byte(jsonStr), &friendDto)

	friendServiceMock := service.FriendServiceMock{}
	friendServiceMock.On("GetCommonFriends", friendDto).Return(friends)
	friendController := FriendController{friendServiceMock}

	r, _ := http.NewRequest("POST", "/friends/get-common-friends-list", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	handler := http.HandlerFunc(friendController.GetCommonFriends)
	handler.ServeHTTP(w, r)

	var actualResult response.Response
	body, _ := ioutil.ReadAll(w.Result().Body)
	json.Unmarshal(body, &actualResult)

	expectedResult := response.Response{Success: true, Friends: friends, Count: len(friends)}

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, expectedResult.Success, actualResult.Success)
	assert.Equal(t, expectedResult.Friends, actualResult.Friends)
	assert.Equal(t, expectedResult.Count, actualResult.Count)
}

func TestGetReceiversList(t *testing.T)  {
	jsonStr := []byte(`{"sender":"a@gmail.com","text":"Hello all, k@gmail.com, a@@gmail.com"}`)
	friends := []string{"b@gmail.com", "k@gmail.com"}

	var senderDto dto.SenderDto
	json.Unmarshal([]byte(jsonStr), &senderDto)

	friendServiceMock := service.FriendServiceMock{}
	friendServiceMock.On("GetReceiversList", senderDto).Return(friends)
	friendController := FriendController{friendServiceMock}

	r, _ := http.NewRequest("POST", "/friends/get-receivers-list", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	handler := http.HandlerFunc(friendController.GetReceiversList)
	handler.ServeHTTP(w, r)

	var actualResult response.Response
	body, _ := ioutil.ReadAll(w.Result().Body)
	json.Unmarshal(body, &actualResult)

	expectedResult := response.Response{Success: true, Friends: friends, Count: len(friends)}

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, expectedResult.Success, actualResult.Success)
	assert.Equal(t, expectedResult.Friends, actualResult.Friends)
	assert.Equal(t, expectedResult.Count, actualResult.Count)
}

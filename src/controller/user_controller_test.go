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

func TestGetAllUsers(t *testing.T)  {
	friends := []string{"a@example.com", "b@example.com", "c@example.com"}

	userServiceMock := service.UserServiceMock{}
	userServiceMock.On("GetAllUsers").Return(friends)
	userController := UserController{userServiceMock}

	r, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	handler := http.HandlerFunc(userController.GetAllUsers)
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

func TestCreateUser(t *testing.T) {
	jsonStr := []byte(`{"email": "a@gmail.com"}`)

	var emailDto dto.EmailDto
	json.Unmarshal([]byte(jsonStr), &emailDto)

	userServiceMock := service.UserServiceMock{}
	userServiceMock.On("CreateUser", emailDto).Return(true)
	userController := UserController{userServiceMock}

	r, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	handler := http.HandlerFunc(userController.CreateUser)
	handler.ServeHTTP(w, r)

	var actualResult response.Success
	body, _ := ioutil.ReadAll(w.Result().Body)
	json.Unmarshal(body, &actualResult)

	expectedResult := response.Success{Success: true}

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, expectedResult.Success, actualResult.Success)
}

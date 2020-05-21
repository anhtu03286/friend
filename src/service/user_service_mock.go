package service

import (
	"github.com/anhtu03286/friend/dto"
	"github.com/anhtu03286/friend/exception"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (u UserServiceMock) GetAllUsers() []string {
	returnVals := u.Called()
	return returnVals.Get(0).([]string)
}

func (u UserServiceMock) CreateUser(emailDto dto.EmailDto) (bool, *exception.Exception) {
	returnVals := u.Called(emailDto)
	return returnVals.Get(0).(bool), nil
}

func (u UserServiceMock) ExistsByEmail(email string) bool {
	returnVals := u.Called(email)
	return returnVals.Get(0).(bool)
}

func (u UserServiceMock) FindUserIdByEmail(email string) (int64, error) {
	returnVals := u.Called(email)
	return returnVals.Get(0).(int64), nil
}

func (u UserServiceMock) FindByIds(ids []int64) []string {
	returnVals := u.Called(ids)
	return returnVals.Get(0).([]string)
}



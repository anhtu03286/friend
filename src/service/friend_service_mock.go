package service

import (
	"github.com/anhtu03286/friend/dto"
	"github.com/anhtu03286/friend/exception"
	"github.com/stretchr/testify/mock"
)

type FriendServiceMock struct {
	mock.Mock
}

func (f FriendServiceMock) CreateFriend(friendDto dto.FriendDto) (bool, *exception.Exception) {
	returnVals := f.Called(friendDto)
	return returnVals.Get(0).(bool), nil
}

func (f FriendServiceMock) CreateSubscribe(requestDto dto.RequestDto) (bool, *exception.Exception) {
	returnVals := f.Called(requestDto)
	return returnVals.Get(0).(bool), nil
}

func (f FriendServiceMock) CreateBlock(requestDto dto.RequestDto) (bool, *exception.Exception) {
	returnVals := f.Called(requestDto)
	return returnVals.Get(0).(bool), nil
}

func (f FriendServiceMock) GetFriendsListByEmail(emailDto dto.EmailDto) ([]string, *exception.Exception) {
	returnVals := f.Called(emailDto)
	return returnVals.Get(0).([]string), nil
}

func (f FriendServiceMock) GetCommonFriends(friendDto dto.FriendDto) ([]string, *exception.Exception) {
	returnVals := f.Called(friendDto)
	return returnVals.Get(0).([]string), nil
}

func (f FriendServiceMock) GetReceiversList(senderDto dto.SenderDto) ([]string, *exception.Exception) {
	returnVals := f.Called(senderDto)
	return returnVals.Get(0).([]string), nil
}

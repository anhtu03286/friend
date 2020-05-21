package service

import (
	"github.com/anhtu03286/friend/entity"
	"github.com/stretchr/testify/mock"

)

type RelationshipServiceMock struct {
	mock.Mock
}

func (r RelationshipServiceMock) CreateRelationship(relationship entity.Relationship) bool {
	returnVals := r.Called(relationship)
	return returnVals.Get(0).(bool)
}

func (r RelationshipServiceMock) FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship {
	returnVals := r.Called(firstEmailId, secondEmailId, status)
	return returnVals.Get(0).([]entity.Relationship)
}

func (r RelationshipServiceMock) FindByEmailIdAndStatus(emailId int64, status []int64) []entity.Relationship {
	returnVals := r.Called(emailId, status)
	return returnVals.Get(0).([]entity.Relationship)
}

func (r RelationshipServiceMock) FindByFirstOrSecondEmailIdAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship {
	returnVals := r.Called(firstEmailId, secondEmailId, status)
	return returnVals.Get(0).([]entity.Relationship)
}

func (r RelationshipServiceMock) FindSubscribersByEmailId(emailId int64) []int64 {
	returnVals := r.Called(emailId)
	return returnVals.Get(0).([]int64)
}

func (r RelationshipServiceMock) IsFriendedOrBlocked(firstEmailId int64, secondEmailId int64) bool {
	returnVals := r.Called(firstEmailId, secondEmailId)
	return returnVals.Get(0).(bool)
}

func (r RelationshipServiceMock) IsSubscribedOrBlocked(firstEmailId int64, secondEmailId int64) bool {
	returnVals := r.Called(firstEmailId, secondEmailId)
	return returnVals.Get(0).(bool)
}

func (r RelationshipServiceMock) IsBlocked(firstEmailId int64, secondEmailId int64) bool {
	returnVals := r.Called(firstEmailId, secondEmailId)
	return returnVals.Get(0).(bool)
}

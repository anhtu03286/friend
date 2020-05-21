package service

import (
	"github.com/anhtu03286/friend/entity"
	"github.com/anhtu03286/friend/enum"
	"github.com/anhtu03286/friend/repository"
)

type IRelationshipService interface {
	CreateRelationship(relationship entity.Relationship) bool
	FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship
	FindByEmailIdAndStatus(emailId int64, status []int64) []entity.Relationship
	FindByFirstOrSecondEmailIdAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship
	FindSubscribersByEmailId(emailId int64) []int64
	IsFriendedOrBlocked(firstEmailId int64, secondEmailId int64) bool
	IsSubscribedOrBlocked(firstEmailId int64, secondEmailId int64) bool
	IsBlocked(firstEmailId int64, secondEmailId int64) bool
}

type RelationshipService struct {
	IRelationshipRepository repository.IRelationshipRepository
}

func (r RelationshipService) CreateRelationship(relationship entity.Relationship) bool {
	return r.IRelationshipRepository.CreateRelationship(relationship)
}

func (r RelationshipService) FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship {
	return r.IRelationshipRepository.FindByTwoEmailIdsAndStatus(firstEmailId, secondEmailId, status)
}

func (r RelationshipService) FindByEmailIdAndStatus(emailId int64, status []int64) []entity.Relationship {
	return r.IRelationshipRepository.FindByEmailIdAndStatus(emailId, status)
}

func (r RelationshipService) FindByFirstOrSecondEmailIdAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship {
	return  r.IRelationshipRepository.FindByFirstOrSecondEmailIdAndStatus(firstEmailId, secondEmailId, status)
}

func (r RelationshipService) FindSubscribersByEmailId(emailId int64) []int64 {
	return r.IRelationshipRepository.FindSubscribersByEmailId(emailId)
}

func (r RelationshipService) IsFriendedOrBlocked(firstEmailId int64, secondEmailId int64) bool {
	relationships := r.FindByTwoEmailIdsAndStatus(firstEmailId, secondEmailId, []int64{enum.FRIEND, enum.BLOCK})
	if relationships != nil && len(relationships) > 0 {
		return true
	}
	return false
}

func (r RelationshipService) IsSubscribedOrBlocked(firstEmailId int64, secondEmailId int64) bool {
	relationships := r.FindByTwoEmailIdsAndStatus(firstEmailId, secondEmailId, []int64{enum.SUBSCRIBE, enum.BLOCK})
	if relationships != nil && len(relationships) > 0 {
		return true
	}
	return false
}

func (r RelationshipService) IsBlocked(firstEmailId int64, secondEmailId int64) bool {
	relationships := r.FindByTwoEmailIdsAndStatus(firstEmailId, secondEmailId, []int64{enum.BLOCK})
	if relationships != nil && len(relationships) > 0 {
		return true
	}
	return false
}
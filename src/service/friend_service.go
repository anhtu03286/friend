package service

import (
	"github.com/anhtu03286/friend/dto"
	"github.com/anhtu03286/friend/entity"
	"github.com/anhtu03286/friend/enum"
	ex "github.com/anhtu03286/friend/exception"
	"github.com/anhtu03286/friend/utils"
	"net/http"
)

type IFriendService interface {
	CreateFriend(friendDto dto.FriendDto) (bool, *ex.Exception)
	CreateSubscribe(requestDto dto.RequestDto) (bool, *ex.Exception)
	CreateBlock(requestDto dto.RequestDto) (bool, *ex.Exception)
	GetFriendsListByEmail(emailDto dto.EmailDto) ([]string, *ex.Exception)
	GetCommonFriends(friendDto dto.FriendDto) ([]string, *ex.Exception)
	GetReceiversList(senderDto dto.SenderDto) ([]string, *ex.Exception)
}

type FriendService struct {
	IRelationshipService IRelationshipService
	IUserService         IUserService
}

func (f FriendService) CreateFriend(friendDto dto.FriendDto) (bool, *ex.Exception) {
	firstEmail := friendDto.Friends[0]
	secondEmail := friendDto.Friends[1]

	if !utils.IsFormatEmail(firstEmail) || !utils.IsFormatEmail(secondEmail) {
		return false, &ex.Exception{Code: http.StatusBadRequest, Message: "Wrong email format!"}
	}

	firstEmailId, err1 := f.IUserService.FindUserIdByEmail(firstEmail)
	if err1 != nil {
		return false, &ex.Exception{Code: http.StatusNotFound, Message: err1.Error()}
	}

	secondEmailId, err2 := f.IUserService.FindUserIdByEmail(secondEmail)
	if err2 != nil {
		return false, &ex.Exception{Code: http.StatusNotFound, Message: err2.Error()}
	}

	// check friend and blocked
	if f.IRelationshipService.IsFriendedOrBlocked(firstEmailId, secondEmailId) {
		return false, &ex.Exception{Code: http.StatusInternalServerError, Message: "Can not make friend!"}
	}

	relationship := entity.Relationship{FirstEmailId: firstEmailId, SecondEmailId: secondEmailId, Status: enum.FRIEND, ActionUserID: firstEmailId}
	result := f.IRelationshipService.CreateRelationship(relationship)
	if result != true {
		return false, &ex.Exception{Code: http.StatusInternalServerError, Message: "Error when create friend!"}
	}
	return true, nil
}

func (f FriendService) CreateSubscribe(requestDto dto.RequestDto) (bool, *ex.Exception) {
	requestor := requestDto.Requestor
	target := requestDto.Target
	if !utils.IsFormatEmail(requestor) || !utils.IsFormatEmail(target) {
		return false, &ex.Exception{Code: http.StatusBadRequest, Message: "Wrong email format!"}
	}

	requestorId, err1 := f.IUserService.FindUserIdByEmail(requestor)
	if err1 != nil {
		return false, &ex.Exception{Code: http.StatusNotFound, Message: err1.Error()}
	}

	targetId, err2 := f.IUserService.FindUserIdByEmail(target)
	if err2 != nil {
		return false, &ex.Exception{Code: http.StatusNotFound, Message: err2.Error()}
	}

	// check requestor has not subscribed target yet
	subscriberEmailIds := f.IRelationshipService.FindSubscribersByEmailId(targetId)
	subscriberEmailIds = utils.RemoveItemFromList(subscriberEmailIds, targetId)

	if subscriberEmailIds != nil && len(subscriberEmailIds) > 0 && utils.Contains(subscriberEmailIds, requestorId) {
		return false, &ex.Exception{Code: http.StatusInternalServerError, Message: "Can not subscribe!"}
	}

	// check both emails have not blocked each other
	if f.IRelationshipService.IsBlocked(requestorId, targetId) {
		return false, &ex.Exception{Code: http.StatusInternalServerError, Message: "Can not subscribe!"}
	}

	relationship := entity.Relationship{FirstEmailId: requestorId, SecondEmailId: targetId, Status: enum.SUBSCRIBE}
	result := f.IRelationshipService.CreateRelationship(relationship)
	if result != true {
		return false, &ex.Exception{Code: http.StatusInternalServerError, Message: "Error when subscribe!"}
	}
	return true, nil
}

func (f FriendService) CreateBlock(requestDto dto.RequestDto) (bool, *ex.Exception) {
	requestor := requestDto.Requestor
	target := requestDto.Target
	if !utils.IsFormatEmail(requestor) || !utils.IsFormatEmail(target) {
		return false, &ex.Exception{Code: http.StatusBadRequest, Message: "Wrong email format!"}
	}

	requestorId, err1 := f.IUserService.FindUserIdByEmail(requestor)
	if err1 != nil {
		return false, &ex.Exception{Code: http.StatusNotFound, Message: err1.Error()}
	}

	targetId, err2 := f.IUserService.FindUserIdByEmail(target)
	if err2 != nil {
		return false, &ex.Exception{Code: http.StatusNotFound, Message: err2.Error()}
	}

	/// check blocked
	if f.IRelationshipService.IsBlocked(requestorId, targetId) {
		return false, &ex.Exception{Code: http.StatusInternalServerError, Message: "Can not subscribe!"}
	}

	relationship := entity.Relationship{FirstEmailId: requestorId, SecondEmailId: targetId, Status: enum.BLOCK}
	result := f.IRelationshipService.CreateRelationship(relationship)
	if result != true {
		return false, &ex.Exception{Code: http.StatusInternalServerError, Message: "Error when subscribe!"}
	}
	return true, nil
}

func (f FriendService) GetFriendsListByEmail(emailDto dto.EmailDto) ([]string, *ex.Exception) {
	var emails []string
	if !utils.IsFormatEmail(emailDto.Email) {
		return nil, &ex.Exception{Code: http.StatusBadRequest, Message: "Wrong email format!"}
	}
	emailId, err1 := f.IUserService.FindUserIdByEmail(emailDto.Email)
	if err1 != nil {
		return nil, &ex.Exception{Code: http.StatusNotFound, Message: err1.Error()}
	}

	// find list relationship by an email and friend status
	relationships := f.IRelationshipService.FindByEmailIdAndStatus(emailId, []int64{enum.FRIEND})
	if relationships != nil {
		// get list email_ids
		emailIds := getEmailIdsFromListRelationships(relationships)
		emailIds = utils.RemoveItemFromList(emailIds, emailId)

		// get list emails by list email_ids
		if emailIds != nil && len(emailIds) > 0 {
			emails = f.IUserService.FindByIds(emailIds)
		}
	}
	return emails, nil
}

func (f FriendService) GetCommonFriends(friendDto dto.FriendDto) ([]string, *ex.Exception) {
	var commonEmails []string

	firstEmail := friendDto.Friends[0]
	secondEmail := friendDto.Friends[1]
	if !utils.IsFormatEmail(firstEmail) || !utils.IsFormatEmail(secondEmail) {
		return nil, &ex.Exception{Code: http.StatusBadRequest, Message: "Wrong email format!"}
	}

	firstEmailId, err1 := f.IUserService.FindUserIdByEmail(firstEmail)
	if err1 != nil {
		return nil, &ex.Exception{Code: http.StatusNotFound, Message: err1.Error()}
	}

	secondEmailId, err2 := f.IUserService.FindUserIdByEmail(secondEmail)
	if err2 != nil {
		return nil, &ex.Exception{Code: http.StatusNotFound, Message: err2.Error()}
	}

	// find list relationship between two email and friend status
	relationships := f.IRelationshipService.FindByFirstOrSecondEmailIdAndStatus(firstEmailId, secondEmailId, []int64{enum.FRIEND})
	if relationships != nil && len(relationships) > 0 {
		mapFirst := make(map[int64]bool)
		var friendsOfFirstEmail []int64

		mapSecond := make(map[int64]bool)
		var friendsOfSecondEmail []int64

		for _, rela := range relationships {
			// get friends of firstEmailId
			if firstEmailId == rela.FirstEmailId || firstEmailId == rela.SecondEmailId {
				if _, ok := mapFirst[rela.FirstEmailId]; !ok {
					mapFirst[rela.FirstEmailId] = true
					friendsOfFirstEmail = append(friendsOfFirstEmail, rela.FirstEmailId)
				}

				if _, ok := mapFirst[rela.SecondEmailId]; !ok {
					mapFirst[rela.SecondEmailId] = true
					friendsOfFirstEmail = append(friendsOfFirstEmail, rela.SecondEmailId)
				}
			}

			// get friends of secondEmailId
			if secondEmailId == rela.FirstEmailId || secondEmailId == rela.SecondEmailId {
				if _, ok := mapSecond[rela.FirstEmailId]; !ok {
					mapSecond[rela.FirstEmailId] = true
					friendsOfSecondEmail = append(friendsOfSecondEmail, rela.FirstEmailId)
				}

				if _, ok := mapSecond[rela.SecondEmailId]; !ok {
					mapSecond[rela.SecondEmailId] = true
					friendsOfSecondEmail = append(friendsOfSecondEmail, rela.SecondEmailId)
				}
			}
		}
		// remove emails in request
		friendsOfFirstEmail = utils.RemoveItemFromList(friendsOfFirstEmail, firstEmailId)
		friendsOfSecondEmail = utils.RemoveItemFromList(friendsOfSecondEmail, secondEmailId)

		// get common emailIds
		var commonEmailIds []int64
		for _, mailId := range friendsOfFirstEmail {
			if utils.Contains(friendsOfSecondEmail, mailId) {
				commonEmailIds = append(commonEmailIds, mailId)
			}
		}

		// get string emails by emailId
		if commonEmailIds != nil && len(commonEmailIds) > 0 {
			commonEmails = f.IUserService.FindByIds(commonEmailIds)
		}
	}
	return commonEmails, nil
}

func (f FriendService) GetReceiversList(senderDto dto.SenderDto) ([]string, *ex.Exception) {
	receiverEmails := []string{}

	senderEmail := senderDto.Sender
	if !utils.IsFormatEmail(senderEmail) {
		return nil, &ex.Exception{Code: http.StatusBadRequest, Message: "Wrong email format!"}
	}

	senderEmailId, err1 := f.IUserService.FindUserIdByEmail(senderEmail)
	if err1 != nil {
		return nil, &ex.Exception{Code: http.StatusNotFound, Message: err1.Error()}
	}

	// get list friend_ids by emailId
	friendRelationships := f.IRelationshipService.FindByEmailIdAndStatus(senderEmailId, []int64{enum.FRIEND})
	friendEmailIds := getEmailIdsFromListRelationships(friendRelationships)
	friendEmailIds = utils.RemoveItemFromList(friendEmailIds, senderEmailId)

	// get subscribers by emailId
	subscriberEmailIds := f.IRelationshipService.FindSubscribersByEmailId(senderEmailId)
	subscriberEmailIds = utils.RemoveItemFromList(subscriberEmailIds, senderEmailId)

	// list receiver emailIds
	receiverEmailIds := friendEmailIds
	for _, mailId := range subscriberEmailIds {
		if !utils.Contains(receiverEmailIds, mailId) {
			receiverEmailIds = append(receiverEmailIds, mailId)
		}
	}

	// get list blocked_ids
	blockedRelationships := f.IRelationshipService.FindByEmailIdAndStatus(senderEmailId, []int64{enum.BLOCK})
	blockedMailIds := getEmailIdsFromListRelationships(blockedRelationships)
	blockedMailIds = utils.RemoveItemFromList(blockedMailIds, senderEmailId)

	// remove blocked in list receiverEmailIds
	for _, mailId := range blockedMailIds {
		receiverEmailIds = utils.RemoveItemFromList(receiverEmailIds, mailId)
	}

	// find email by list receiverEmailIds
	receiverEmails = f.IUserService.FindByIds(receiverEmailIds)

	// get mentionedEmail from text of sender
	if &senderDto.Text != nil {
		mentionEmails := utils.GetEmailFromText(senderDto.Text)
		if &mentionEmails != nil && len(mentionEmails) > 0 {
			for _, email := range mentionEmails {
				receiverEmails = append(receiverEmails, email)
			}
		}
	}
	return receiverEmails, nil
}

func getEmailIdsFromListRelationships(relationships []entity.Relationship) []int64 {
	keys := make(map[int64]bool)
	set := []int64{}
	if relationships != nil && len(relationships) > 0 {
		for _, rela := range relationships {
			if _, ok := keys[rela.FirstEmailId]; !ok {
				keys[rela.FirstEmailId] = true
				set = append(set, rela.FirstEmailId)
			}

			if _, ok := keys[rela.SecondEmailId]; !ok {
				keys[rela.SecondEmailId] = true
				set = append(set, rela.SecondEmailId)
			}
		}
	}
	return set
}
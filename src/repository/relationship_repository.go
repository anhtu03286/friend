package repository

import (
	"database/sql"
	"fmt"
	"github.com/anhtu03286/friend/entity"
	"strconv"
	"strings"
)

type IRelationshipRepository interface {
	CreateRelationship(relationship entity.Relationship) bool
	FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship
	FindByEmailIdAndStatus(emailId int64, status []int64) []entity.Relationship
	FindByFirstOrSecondEmailIdAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship
	FindSubscribersByEmailId(emailId int64) []int64
}

type RelationshipRepository struct {
	DB *sql.DB
}


func (r RelationshipRepository) CreateRelationship(relationship entity.Relationship) bool {
	query, err := r.DB.Prepare(`insert into relationships (first_email_id, second_email_id, status, action_user_id) values (?, ?, ?, ?);`)
	if err != nil {
		return false
	}
	_, _ = query.Exec(relationship.FirstEmailId, relationship.SecondEmailId, relationship.Status, relationship.ActionUserID)
	return true
}

func (r RelationshipRepository) FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship {
	strStatusIds := make([]string, len(status))
	for i, id := range status {
		strStatusIds[i] = strconv.FormatInt(id, 10)
	}

	stmt := `select r.*	from relationships r where r.first_email_id in (%s, %s)	and r.second_email_id in (%s, %s) and r.status in (%s);`
	query := fmt.Sprintf(
		stmt,
		strconv.FormatInt(firstEmailId, 10),
		strconv.FormatInt(secondEmailId, 10),
		strconv.FormatInt(firstEmailId, 10),
		strconv.FormatInt(secondEmailId, 10),
		strings.Join(strStatusIds, ","))

	results, err := r.DB.Query(query)
	if err != nil {
		panic(err)
	}

	var relationships []entity.Relationship
	for results.Next() {
		var id, firstEmailId, secondEmailId, status int64
		_ = results.Scan(&id, &firstEmailId, &secondEmailId, &status)
		relationship := entity.Relationship{Id: id, FirstEmailId: firstEmailId, SecondEmailId: secondEmailId, Status: status}
		relationships = append(relationships, relationship)
	}
	return relationships
}

func (r RelationshipRepository) FindByEmailIdAndStatus(emailId int64, status []int64) []entity.Relationship {
	strStatusIds := make([]string, len(status))
	for i, id := range status {
		strStatusIds[i] = strconv.FormatInt(id, 10)
	}

	stmt := `select r.* from relationships r where (r.first_email_id = %s or r.second_email_id = %s) and r.status in (%s);`
	query := fmt.Sprintf(
		stmt,
		strconv.FormatInt(emailId, 10),
		strconv.FormatInt(emailId, 10),
		strings.Join(strStatusIds, ","))

	results, err := r.DB.Query(query)
	if err != nil {
		panic(err)
	}

	var relationships []entity.Relationship
	for results.Next() {
		var id, firstEmailId, secondEmailId, status, actionUserId int64
		results.Scan(&id, &firstEmailId, &secondEmailId, &status, &actionUserId)
		relationship := entity.Relationship{Id: id, FirstEmailId: firstEmailId, SecondEmailId: secondEmailId, Status: status, ActionUserID: actionUserId}
		relationships = append(relationships, relationship)
	}
	return relationships
}

func (r RelationshipRepository) FindByFirstOrSecondEmailIdAndStatus(firstEmailId int64, secondEmailId int64, status []int64) []entity.Relationship {
	strStatusIds := make([]string, len(status))
	for i, id := range status {
		strStatusIds[i] = strconv.FormatInt(id, 10)
	}

	stmt := `select r.* from relationships r where r.first_email_id in (%s, %s) or r.second_email_id in (%s, %s) and r.status in (%s);`
	query := fmt.Sprintf(
		stmt,
		strconv.FormatInt(firstEmailId, 10),
		strconv.FormatInt(secondEmailId, 10),
		strconv.FormatInt(firstEmailId, 10),
		strconv.FormatInt(secondEmailId, 10),
		strings.Join(strStatusIds, ","))

	results, err := r.DB.Query(query)
	if err != nil {
		panic(err)
	}

	var relationships []entity.Relationship
	for results.Next() {
		var id, firstEmailId, secondEmailId, status, actionUserId int64
		_ = results.Scan(&id, &firstEmailId, &secondEmailId, &status, &actionUserId)
		relationship := entity.Relationship{Id: id, FirstEmailId: firstEmailId, SecondEmailId: secondEmailId, Status: status, ActionUserID: actionUserId}
		relationships = append(relationships, relationship)
	}
	return relationships
}

func (r RelationshipRepository) FindSubscribersByEmailId(emailId int64) []int64 {
	query := `select r.first_email_id from relationships r where r.second_email_id = ? and r.status = 1;`
	results, err := r.DB.Query(query, emailId)
	if err != nil {
		panic(err)
	}

	var emailIds []int64
	for results.Next() {
		var id int64
		_ = results.Scan(&id)
		emailIds = append(emailIds, id)
	}
	return emailIds
}

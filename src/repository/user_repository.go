package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type IUserRepository interface {
	GetAllUsers() []string
	CreateUser(email string) bool
	ExistsByEmail(email string) bool
	FindUserIdByEmail(email string) int64
	FindByIds(ids []int64) []string
}

type UserRepository struct {
	DB *sql.DB
}

// use for testing
func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{DB: db}
}

func (u UserRepository) GetAllUsers() []string {
	result, err := u.DB.Query("select email from users")
	if err != nil {
		panic(err.Error())
	}

	var emails []string
	for result.Next() {
		var email string
		err = result.Scan(&email)
		if err != nil {
			panic(err.Error())
		}
		emails = append(emails, email)
	}
	return emails
}

func (u UserRepository) CreateUser(email string) bool {
	query, err := u.DB.Prepare(`insert into users (email) values (?)`)
	if err != nil {
		return false
	}
	query.Exec(email)
	return true
}

func (u UserRepository) ExistsByEmail(email string) bool {
	var id int
	err := u.DB.QueryRow(`select id from users where email=?`, email).Scan(&id)
	if err != nil {
		return false
	}
	return true
}

func (u UserRepository) FindUserIdByEmail(email string) int64 {
	var id int64
	err := u.DB.QueryRow("select id from users where email=?", email).Scan(&id)
	if err != nil {
		return -1
	}
	return id
}

func (u UserRepository) FindByIds(ids []int64) []string {
	strIds := make([]string, len(ids))
	for i, id := range ids {
		strIds[i] = strconv.FormatInt(id, 10)
	}

	stmt := `select u.email	from users u where u.id in (%s);`
	query := fmt.Sprintf(stmt, strings.Join(strIds, ","))
	results, err := u.DB.Query(query)
	if err != nil {
		panic(err)
	}

	emails := []string{}
	for results.Next() {
		var email string
		results.Scan(&email)
		emails = append(emails, email)
	}
	return emails
}
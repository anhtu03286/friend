package service

import (
	"errors"
	"github.com/anhtu03286/friend/dto"
	exp "github.com/anhtu03286/friend/exception"
	"github.com/anhtu03286/friend/repository"
	"github.com/anhtu03286/friend/utils"
	"net/http"
)

type IUserService interface {
	GetAllUsers() []string
	CreateUser(emailDto dto.EmailDto) (bool, *exp.Exception)
	ExistsByEmail(email string) bool
	FindUserIdByEmail(email string) (int64, error)
	FindByIds(ids []int64) []string
}

type UserService struct {
	IUserRepository repository.IUserRepository
}

func (f UserService) GetAllUsers() []string {
	return f.IUserRepository.GetAllUsers()
}

func (f UserService) CreateUser(emailDto dto.EmailDto) (bool, *exp.Exception) {
	if !utils.IsFormatEmail(emailDto.Email) {
		return false, &exp.Exception{Code: http.StatusBadRequest, Message: "Wrong email format!"}
	}

	if f.IUserRepository.ExistsByEmail(emailDto.Email) {
		return false, &exp.Exception{Code: http.StatusInternalServerError, Message: "Email already existed!"}
	}

	user := f.IUserRepository.CreateUser(emailDto.Email)
	if user != true {
		return false, &exp.Exception{Code: http.StatusInternalServerError, Message: "Error when create user!"}
	}
	return true, nil
}

func (f UserService) ExistsByEmail(email string) bool {
	return f.IUserRepository.ExistsByEmail(email)
}

func (f UserService) FindUserIdByEmail(email string) (int64, error) {
	id := f.IUserRepository.FindUserIdByEmail(email)
	if id == -1 {
		return -1, errors.New("User not found!")
	}
	return id, nil
}

func (f UserService) FindByIds(ids []int64) []string {
	return f.IUserRepository.FindByIds(ids)
}
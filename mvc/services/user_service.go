package services

import (
	"github.com/odealidj/basic-go-microservices/mvc/domain"
	"github.com/odealidj/basic-go-microservices/mvc/utils"
)

var (
	UsersService usersService
)

type usersService struct {
	
}


func (u *usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError )  {
	user, err := domain.UserDao.GetUser(userId)
	if err!=nil{
		return nil, err
	}
	return user,nil
}

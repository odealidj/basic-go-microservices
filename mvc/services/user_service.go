package services

import (
	"github.com/odealidj/basic-go-microservices/mvc/domain"
	"github.com/odealidj/basic-go-microservices/mvc/utils"
)



func GetUser(userId int64) (*domain.User, *utils.ApplicationError )  {
	return domain.GetUser(userId)
}

package domain

import (
	"fmt"
	"github.com/odealidj/basic-go-microservices/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Ali", LastName: "Laode", Email: "odealidj.go@gmail.com"},
	}

	UserDao userDaoInterface
)

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {
	
}

func init()  {
	UserDao = &userDao{}
}

func (u *userDao)GetUser(userId int64) (*User, *utils.ApplicationError)  {
	log.Println("we're accessing the database")
	if user:= users[userId]; user !=nil {
		return user, nil

	}
	return nil,
		&utils.ApplicationError{
			Message: fmt.Sprintf("user %v does not exists", userId),
			StatusCode: http.StatusNotFound,
			Code: "not_found",
		}


}

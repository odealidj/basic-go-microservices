package services

import (
	"github.com/odealidj/basic-go-microservices/mvc/domain"
	"github.com/odealidj/basic-go-microservices/mvc/utils"
	"net/http"
)

type itemsService struct {

}

var (
	ItemsService itemsService
)

func (i *itemsService)GetItem(itemId string) (*domain.Item, *utils.ApplicationError )  {
	return nil, &utils.ApplicationError{
		Message: "Implement me",
		StatusCode: http.StatusInternalServerError,
	}
}

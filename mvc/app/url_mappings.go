package app

import (
	"github.com/odealidj/basic-go-microservices/mvc/controllers"
)

func mapUrls()  {
	//http.HandleFunc("/users", controllers.GetUser)

	router.GET("/users/:user_id", controllers.GetUser)


}
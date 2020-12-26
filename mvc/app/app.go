package app

import (
	"github.com/odealidj/basic-go-microservices/mvc/controllers"
	"net/http"
)

func StartApp()  {
	http.HandleFunc("/users", controllers.GetUser)
	if err:= http.ListenAndServe(":9000",nil); err!=nil {
		panic(err)
	}
}

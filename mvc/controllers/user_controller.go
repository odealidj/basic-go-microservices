package controllers

import (
	"encoding/json"
	"github.com/odealidj/basic-go-microservices/mvc/services"
	"github.com/odealidj/basic-go-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request  )  {
	userIdParam:=req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam,10,64)
	if err!=nil {
		apiErr:= &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		jsonValue,_ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	user,apiErr :=  services.GetUser(userId)
	if apiErr!=nil {
		jsonValue,_ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	//return user to client
	jsonValue,_ := json.Marshal(user)
	res.Write(jsonValue)


}

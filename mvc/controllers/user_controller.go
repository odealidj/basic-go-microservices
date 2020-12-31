package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/odealidj/basic-go-microservices/mvc/services"
	"github.com/odealidj/basic-go-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context )  {
	userId, err := strconv.ParseInt(c.Param("user_id"),10,64)
	if err!=nil {
		apiErr:= &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}



	user,apiErr :=  services.UsersService.GetUser(userId)
	if apiErr!=nil {
		utils.RespondError(c, apiErr)
		return
	}

	//return user to client
	utils.Respond(c, http.StatusOK, user)


}

/*
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



	user,apiErr :=  services.UsersService.GetUser(userId)
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
*/
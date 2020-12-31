package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine

)

func init()  {
	router = gin.Default()


}

func StartApp()  {
	mapUrls()

	//if err:= http.ListenAndServe(":9000",nil); err!=nil {

	if err:= router.Run(":9000"); err!=nil {
		panic(err)
	}

}

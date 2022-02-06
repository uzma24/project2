package handler

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/getFrequency", serviceCall)
	return router

}

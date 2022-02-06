package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project2/internal"
	"github.com/project2/internal/utils"
)

func serviceCall(c *gin.Context) {
	resp, err := internal.GetFrequency(c)
	if err != nil {
		utils.HandleError(c, err)
		log.Printf("error occured while calling to GET GetFrequency: %v", err)
		return
	}
	utils.SuccessResponse(c, gin.H{
		"status":  http.StatusOK,
		"message": "frequency received successfully",
		"error":   false,
		"data":    resp,
	})
}

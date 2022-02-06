package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BadRequest struct {
	ErrMessage string
}

func (e *BadRequest) Error() string {
	return fmt.Sprintf("%v: BadRequestError", e.ErrMessage)
}

type InternalServerError struct {
	ErrMessage string
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("%v: InternalServerError", e.ErrMessage)
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, data)
}

func HandleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *BadRequest:
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"status":  http.StatusBadRequest,
			"message": e.ErrMessage,
		})
		return
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"status":  http.StatusInternalServerError,
			"message": e.Error(),
		})
		return
	}
}

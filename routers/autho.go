package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	m "ymo.me/sbum-end/models"
)

func login(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestUser m.User

	if err := context.ShouldBindJSON(&requestUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "Request to user login",
	})
}

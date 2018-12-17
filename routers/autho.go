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
	if requestUser.Pass != "123" {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"authUser": requestUser,
	})
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/auth"
	"server/database"
	"server/models"
)

type LoginRequest struct {
	//Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenRequest struct {
	//Email    string `json:"email"`
	Username string `json:"username"`
}

func Login(context *gin.Context) {
	var request LoginRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	//check if username exists and password is correct
	record := database.Instance.Where("username = ?", request.Username).Find(&user)
	//record := database.Instance.Where("username = ?", request.Username).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login success", "token": tokenString})
	return
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	//check if username exists and password is correct
	record := database.Instance.Where("username = ?", request.Username).Find(&user)
	//record := database.Instance.Where("username = ?", request.Username).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	//credentialError := user.CheckPassword(request.Password)
	//if credentialError != nil {
	//	context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
	//	context.Abort()
	//	return
	//}

	tokenString, err := auth.GenerateJWT(user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Generate token success", "token": tokenString})
	return
}

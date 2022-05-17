package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/database"
	"server/models"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "username": user.Username})
	return
}

func GetUsers(context *gin.Context) {
	var user models.User

	if result := database.Instance.Find(&user); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"userId": user.ID, "name": user.Name, "username": user.Username})
	return
}

func Logout(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Logout success"})
	context.Redirect(http.StatusMovedPermanently, "/")
	return
}

//func Login(context *gin.Context) {
//	var user models.User
//	//var request TokenRequest
//
//	if err := context.ShouldBindJSON(&user); err != nil {
//		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		context.Abort()
//		return
//	}
//	//if err := user.HashPassword(user.Password); err != nil {
//	//	context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//	//	context.Abort()
//	//	return
//	//}
//
//	result := database.Instance.Where("username = ?", user.Username).Find(&user)
//
//	if result.Error != nil {
//		context.JSON(http.StatusBadRequest, gin.H{"message": "Problem logging into your account"})
//		context.Abort()
//		return
//	}
//
//	if user.Username == "" {
//		context.JSON(http.StatusNotFound, gin.H{"message": "User account was not found"})
//		context.Abort()
//		return
//	}
//
//	jwtToken := GenerateToken
//	context.JSON(200, gin.H{"message": "Log in success", "token": jwtToken})
//}

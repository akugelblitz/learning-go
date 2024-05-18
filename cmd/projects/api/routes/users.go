package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/api/models"
	"example.com/api/utils"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save to db"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.ValidateCreds()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, int(user.ID))
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not authenticate user"},
		)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

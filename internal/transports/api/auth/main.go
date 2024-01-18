package auth

import (
	"cms/internal/models/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Auth gin.HandlerFunc = func(c *gin.Context) {
	var input auth.UserModel

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rightPass, err := input.CheckUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !rightPass {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le mot de passe ne correspond pas"})
		return
	}

	token.GenerateToken
}

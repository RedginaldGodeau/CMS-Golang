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
}

package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izabela-k/lesfinder/utils"
)

type SignUpData struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Password2 string `json:"password2" binding:"required,eqfield=Password,min=6"`
}

func SignUpHandler(c *gin.Context) {
	var data SignUpData

	if err := c.ShouldBindJSON(&data); err != nil {
		errors := utils.HandleValidatorErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	user := new(User)
	user.CreateUser(data.Email, data.Password)

	c.JSON(200, gin.H{
		"message": "OK",
	})
}


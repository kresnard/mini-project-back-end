package app

import (
	"mpbe/input"
	"mpbe/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	service service.UserService
}

func (ch *UserHandlers) createUser(c *gin.Context) {
	var input input.UserInput
	err := c.ShouldBindJSON(&input)
	users, _ := ch.service.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, users)
}

package app

import (
	"fmt"
	"mpbe/auth"
	"mpbe/domain"
	"mpbe/input"
	"mpbe/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userDTO struct {
	ID    int    `json:"user_id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func FormatuserDTO(user domain.Users, token string) userDTO {
	userDTO := userDTO{}
	userDTO.ID = user.ID
	userDTO.Email = user.Email
	userDTO.Role = user.Role
	userDTO.Token = token
	return userDTO
}

type UserHandlers struct {
	service     service.UserService
	authService auth.AuthService
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

func (ch *UserHandlers) loginUser(c *gin.Context) {
	var input input.Login
	err := c.ShouldBindJSON(&input)
	users, _ := ch.service.LoginUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	token, err := ch.authService.GenerateToken(users.ID)
	if err != nil {
		fmt.Println("TEST", token)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	userDTO := FormatuserDTO(users, token)
	c.JSON(http.StatusOK, userDTO)
}

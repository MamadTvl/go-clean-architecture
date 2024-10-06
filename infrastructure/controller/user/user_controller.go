package user_controller

import (
	user_interactor "clean-architecture/use-case/user"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	saveUserUseCase user_interactor.SaveUserUseCase
}

func NewController(saveUserUseCase *user_interactor.SaveUserUseCase) *Controller {
	return &Controller{
		saveUserUseCase: *saveUserUseCase,
	}
}

type CreateUserDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *Controller) CreateUser(c *gin.Context) {
	var dto CreateUserDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"message": "validation error", "details": err})
		return
	}

	user, err := u.saveUserUseCase.SaveUser(dto.Username, dto.Password)

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(400, gin.H{"message": "username already exists"})
		}
		c.JSON(400, gin.H{"message": "something went wrong", "details": err})
		return
	}

	c.JSON(200, gin.H{"data": "user created", "username": user.UserName})
}

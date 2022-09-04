package ucon

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tnakamura512/go-modular-monolith-base/modules/user/umodel"
)

type UserController struct {
	userModel *umodel.UserModel
}

func NewUserController(userModel *umodel.UserModel) *UserController {
	controller := new(UserController)
	controller.userModel = userModel
	return controller
}

func (u *UserController) AddUser(c *gin.Context) {
	request := AddUserRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response_code": "ER",
			"message":       err.Error(),
		})
		return
	}

	toekn, err := u.userModel.AddUser(request.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response_code": "ER",
			"message":       err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":         toekn,
		"response_code": "OK",
	})
}

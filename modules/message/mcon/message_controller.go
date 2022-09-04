package mcon

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tnakamura512/go-modular-monolith-base/modules/message/mmodel"
)

type MessageController struct {
	messageModel *mmodel.MessageModel
}

func NewMessageController(messageModel *mmodel.MessageModel) *MessageController {
	controller := new(MessageController)
	controller.messageModel = messageModel
	return controller
}

func (m *MessageController) SendMessage(c *gin.Context) {
	request := SendMessageRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response_code": "ER",
			"message":       err.Error(),
		})
		return
	}

	err := m.messageModel.DoSendMessage(request.Token, request.Message, request.SendIds)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response_code": "ER",
			"message":       err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response_code": "OK",
	})
}

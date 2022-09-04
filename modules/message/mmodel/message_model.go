package mmodel

import (
	"fmt"

	"github.com/tnakamura512/go-modular-monolith-base/modules"
)

type MessageModel struct {
	repo            MessageRepositroyInterface
	moduleApiBridge *modules.ModuleApiBridge
}

func NewMessageModel(repo MessageRepositroyInterface,
	moduleApiBridge *modules.ModuleApiBridge) *MessageModel {
	model := new(MessageModel)
	model.repo = repo
	model.moduleApiBridge = moduleApiBridge
	return model
}

func (m *MessageModel) DoSendMessage(apiToken, message string, toUserIds []int) error {
	senderId, err := m.moduleApiBridge.AuthUser(apiToken)
	if err != nil {
		return err
	}

	return m.SendMessage(senderId, toUserIds, message)
}

func (m *MessageModel) SendMessage(senderId int, userIds []int, message string) error {
	fmt.Printf("sending message from %d to %d", userIds, userIds)
	// メッセージPush
	//　Repository経由で送信履歴保存等...
	return m.repo.SaveHistory(message, senderId, userIds)
}

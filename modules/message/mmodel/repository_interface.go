package mmodel

type MessageRepositroyInterface interface {
	SaveHistory(message string, senderId int, toUserIds []int) error
}

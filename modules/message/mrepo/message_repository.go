package mrepo

import "github.com/jmoiron/sqlx"

type MessageRepository struct {
	db *sqlx.DB
}

func NewMessageRepo(db *sqlx.DB) *MessageRepository {
	repo := new(MessageRepository)
	repo.db = db
	return repo
}

func (m *MessageRepository) SaveHistory(message string, senderId int, toUserIds []int) error {
	panic("not implemented")
}

package urepo

import "github.com/jmoiron/sqlx"

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepository {
	repo := new(UserRepository)
	repo.db = db
	return repo
}

func (u *UserRepository) GetIdByToken(token string) (int, error) {
	// dummy impl
	// return 100, nil
	panic("not implemented")

}

func (u *UserRepository) SaveUser(token, name string) error {
	panic("not implemented")
}

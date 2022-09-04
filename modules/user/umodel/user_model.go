package umodel

import "github.com/tnakamura512/go-modular-monolith-base/modules"

type UserModel struct {
	userRepo        UserRepositroyInterface
	moduleApiBridge *modules.ModuleApiBridge
}

func NewUserModel(userRepo UserRepositroyInterface,
	moduleApiBridge *modules.ModuleApiBridge) *UserModel {
	model := new(UserModel)
	model.userRepo = userRepo
	model.moduleApiBridge = moduleApiBridge
	return model
}

// 他のモジュールに公開してるAPI
func (u *UserModel) AuthUser(apiToken string) (int, error) {
	userId, err := u.userRepo.GetIdByToken(apiToken)
	if err != nil {
		return -1, err
	}

	return userId, err
}

// Token発行、UserDBに保存、Token返却みたいな
func (u *UserModel) AddUser(name string) (string, error) {
	token := "GeneratedToken"

	err := u.userRepo.SaveUser(token, name)
	if err != nil {
		return "", err
	}
	return token, err
}

package modules

// モジュール間API
type UserModelApi interface {
	// 認証が成功したらUserIDを返す
	AuthUser(apiToken string) (int, error)
}

type MessageModelApi interface {
	// 指定したユーザ一覧にメッセージを送信する
	SendMessage(senderId int, userIds []int, message string) error
}

type ModuleApiBridge struct {
	userModel UserModelApi
	msgModel  MessageModelApi
}

func NewModuleApiBridge() *ModuleApiBridge {
	return new(ModuleApiBridge)
}

func (m *ModuleApiBridge) RegisterUserModel(model UserModelApi) {
	m.userModel = model
}

func (m *ModuleApiBridge) RegisterMessageModel(model MessageModelApi) {
	m.msgModel = model
}

func (m *ModuleApiBridge) AuthUser(apiToken string) (int, error) {
	return m.userModel.AuthUser(apiToken)
}

func (m *ModuleApiBridge) SendMessage(senderId int, userIds []int, message string) error {
	return m.msgModel.SendMessage(senderId, userIds, message)
}

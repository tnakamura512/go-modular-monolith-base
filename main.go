package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/tnakamura512/go-modular-monolith-base/database"
	"github.com/tnakamura512/go-modular-monolith-base/modules"
	"github.com/tnakamura512/go-modular-monolith-base/modules/message/mcon"
	"github.com/tnakamura512/go-modular-monolith-base/modules/message/mmodel"
	"github.com/tnakamura512/go-modular-monolith-base/modules/message/mrepo"
	"github.com/tnakamura512/go-modular-monolith-base/modules/user/ucon"
	"github.com/tnakamura512/go-modular-monolith-base/modules/user/umodel"
	"github.com/tnakamura512/go-modular-monolith-base/modules/user/urepo"
)

func main() {
	fmt.Println("module server start")

	fmt.Println("DB準備")
	userDb := database.Connection.OpenUserDb()
	defer func() {
		userDb.Close()
	}()
	msgDb := database.Connection.OpenMessageDb()
	defer func() {
		msgDb.Close()
	}()
	// MySQLの例だが、ランキングなのでRedisとかにすることも可能
	stepRankDb := database.Connection.OpenStepRankeDb()
	defer func() {
		stepRankDb.Close()
	}()

	ginEngine := gin.New()
	moduleApiBridge := modules.NewModuleApiBridge()

	fmt.Println("各モジュール準備、依存の注入など")
	setupUserModule(userDb, ginEngine, moduleApiBridge)
	setupMessageModule(msgDb, ginEngine, moduleApiBridge)
	setupStepRankModule(stepRankDb, ginEngine, moduleApiBridge)
}

func setupUserModule(
	db *sqlx.DB,
	ginEngine *gin.Engine,
	moduleApiBridge *modules.ModuleApiBridge,
) {
	userRepo := urepo.NewUserRepo(db)
	userModel := umodel.NewUserModel(userRepo, moduleApiBridge)
	moduleApiBridge.RegisterUserModel(userModel)

	userController := ucon.NewUserController(userModel)
	ginEngine.POST("/user/add", userController.AddUser)
}

func setupMessageModule(
	db *sqlx.DB,
	ginEngine *gin.Engine,
	moduleApiBridge *modules.ModuleApiBridge,
) {
	msgRepo := mrepo.NewMessageRepo(db)
	msgModel := mmodel.NewMessageModel(msgRepo, moduleApiBridge)
	moduleApiBridge.RegisterMessageModel(msgModel)

	msgController := mcon.NewMessageController(msgModel)
	ginEngine.POST("/message/send", msgController.SendMessage)
}

func setupStepRankModule(
	db *sqlx.DB,
	ginEngine *gin.Engine,
	moduleApiBridge *modules.ModuleApiBridge,
) {
	fmt.Println("dummy setup")
}

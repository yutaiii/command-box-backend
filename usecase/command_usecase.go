package usecase

import (
	"github.com/yutaiii/command-box-backend/entity"
	"github.com/yutaiii/command-box-backend/store"
	"github.com/yutaiii/command-box-backend/tool/util"
)

func GetCommands() (*[]entity.Command, error) {
	db := util.GetDBConnection()
	return store.GetAll(db)
}

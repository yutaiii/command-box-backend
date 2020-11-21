package store

import (
	"github.com/yutaiii/command-box-backend/entity"

	"gorm.io/gorm"
)

func GetAll(db *gorm.DB) (*[]entity.Command, error) {
	commands := &[]entity.Command{}
	err := db.Find(commands).Error
	if err != nil {
		return nil, err
	}
	return commands, nil
}

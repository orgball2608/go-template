package mysql

import (
	"elearning/infra/mysql/model"
	"gorm.io/gorm"
)

func addUserTable(db *gorm.DB) error {
	var err error
	if err = db.AutoMigrate(&model.User{}); err != nil {
		return err
	}
	return nil
}

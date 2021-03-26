package services

import (
	"github.com/maus/basic/app/models"
	"github.com/maus/basic/config"
)

func InitTables() error {
	db := config.DB
	if db.Migrator().HasTable(&models.User{}) == false {
		db.Migrator().CreateTable(&models.User{})
		var item models.User
		item.Name = "test role"
		item.DisplayName = "test"
		item.Email = "test@test.com"
		item.EmailIsVerified = true
		item.Tel = "13800138000"
		item.Password = "secret"
		item.Status = true
		db.Create(&item)
	}

	//if db.Migrator().HasColumn(&models.Category{}, "SelectId") == false {
	//	db.Migrator().AddColumn(&models.Category{}, "SelectId")
	//}
	//if db.Migrator().HasTable(&models.Laji{}) == false {
	//	db.Migrator().CreateTable(&models.Laji{})
	//}
	return nil

}

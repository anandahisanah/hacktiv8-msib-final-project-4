package database

import (
	"errors"
	"hacktiv8-msib-final-project-4/config"
	"hacktiv8-msib-final-project-4/entity"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"log"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func seedAdmin() {
	admin := &entity.User{
		FullName: "admin",
		Email:    "admin@hacktiv8.com",
		Password: "admin123",
		Role:     "admin",
		Balance:  0,
	}
	errs.CheckErr(admin.HashPassword())

	errs.CheckErr(db.Create(admin).Error)

	log.Println("Admin account seed success!")
}

func init() {
	db, err = gorm.Open(config.GetDBConfig())
	errs.CheckErr(err)

	errs.CheckErr(db.AutoMigrate(&entity.User{}))

	if db.Migrator().HasTable(&entity.User{}) {
		if err := db.First(&entity.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seedAdmin()
		}
	}

	log.Println("Connected to DB!")
}

func GetPostgresInstance() *gorm.DB {
	return db
}

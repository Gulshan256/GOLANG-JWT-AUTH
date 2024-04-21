package initializers

import (
	"github.com/Gulshan256/GOLANG-JWT-AUTH/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {

	var err error
	DB, err = gorm.Open(postgres.Open("postgres://postgres:123@localhost:5432/go_tut"), &gorm.Config{})
	if err != nil {
		panic(" to connect database failed")
	}
	DB.AutoMigrate(&models.AuthUser{})

}

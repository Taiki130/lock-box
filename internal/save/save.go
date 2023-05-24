package save

import (
	"log"

	"gorm.io/gorm"
)

func SavePassword(db *gorm.DB, username, password, description string) {
	p := Password{
		Username:    username,
		Password:    password,
		Description: description,
	}

	err := db.Create(&p).Error
	if err != nil {
		log.Fatal(err)
	}
}

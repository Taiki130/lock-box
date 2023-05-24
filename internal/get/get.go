package get

import (
	"log"

	"gorm.io/gorm"
)

func GetPassword(db *gorm.DB, username string) {
	var password Password
	err := db.Where("username = ?", username).First(&password).Error
	if err != nil {
		log.Fatal(err)
	}

	return password
}

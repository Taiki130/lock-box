package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Taiki130/lock-box/internal/get"
	"github.com/Taiki130/lock-box/internal/save"
)

type Password struct {
	gorm.Model
	Username    string
	Password    string
	Description string
}

func main() {

	dsn := "root:password@tcp(localhost:3306)/lockbox?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.AutoMigrate(&Password{})
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Please specify a subcommand")
		os.Exit(1)
	}

	subcommand := flag.Arg(0)

	switch subcommand {
	case "save":
		user := flag.String("user", "", "Specify user name to save")
		pass := flag.String("pass", "", "Specify password to save")
		description := flag.String("description", "", "Specify description to save")
		flag.Parse()

		if *user == "" {
			fmt.Println("Please specify user name using the -user flag")
			return
		}

		if *pass == "" {
			fmt.Println("Please specify password using the -pass flag")
			return
		}

		if *description == "" {
			fmt.Println("Please specify password using the -description flag")
			return
		}

		save.SavePassword(db, *user, *pass, *description)

	case "get":
		user := flag.String("user", "", "Specify your user name")
		flag.Parse()

		if *user == "" {
			fmt.Println("Please specify your user name using the -user flag")
			return
		}

		password := get.GetPassword(db, *user)

		log.Printf("Username: %s, Password: %s, Description: %s\n", password.Username, password.Password, password.Description)

	default:
		fmt.Println("無効なサブコマンドです。")
		os.Exit(1)
	}

}

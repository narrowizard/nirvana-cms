package services

import (
	"go-cms/models"

	"github.com/jinzhu/gorm"
	"github.com/kdada/tinygo/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func init() {
	var cfg, err = config.NewConfig("ini", "./config/db.cfg")
	checkErr(err)
	connString, err := cfg.GlobalSection().String("ConnectionString")
	checkErr(err)
	db, err = gorm.Open("mysql", connString)
	checkErr(err)
	createDatabase()
}

func createDatabase() {
	db.AutoMigrate(&models.User{})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

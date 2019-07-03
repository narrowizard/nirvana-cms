package services

import (
	"github.com/narrowizard/nirvana-cms/models"

	"github.com/jinzhu/gorm"
	"github.com/kdada/tinygo/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB
var configInfo = &models.ConfigInfo{}

func init() {
	var cfg, err = config.NewConfig("ini", "./config/config.cfg")
	checkErr(err)
	connString, err := cfg.GlobalSection().String("ConnectionString")
	checkErr(err)
	db, err = gorm.Open("mysql", connString)
	checkErr(err)
	port, err := cfg.GlobalSection().Int("Port")
	checkErr(err)
	configInfo.Port = uint16(port)
	createDatabase()
}

func ConfigInfo() *models.ConfigInfo {
	return configInfo
}

func createDatabase() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Menu{})
	db.AutoMigrate(&models.UserMenu{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.RoleMenu{})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

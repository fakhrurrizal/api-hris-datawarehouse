package config

import (
	"fmt"
	"hris-datawarehouse/app/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBS = make(map[string]*gorm.DB)

// Database Initialization
func Database() *gorm.DB {
	cfg := LoadConfig()
	host := cfg.DatabaseHost
	user := cfg.DatabaseUsername
	password := cfg.DatabasePassword
	name := cfg.DatabaseName
	port := cfg.DatabasePort

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		user, password, host, port, name)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to MySQL: " + err.Error())
	}

	err = DB.AutoMigrate(
		&models.GlobalUser{},
	)
	if err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	} else {
		fmt.Println("Auto migration success ...")

	}

	fmt.Println("Connected to MySQL Database:", name)
	return DB
}

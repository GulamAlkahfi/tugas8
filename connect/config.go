package config

import (
  "tugas8/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)


var DB *gorm.DB
func InitDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/sensus1?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
        panic("failed to connect database")
    }
    InitMigration()
}
func InitMigration(){
  	DB.AutoMigrate(&models.Pekerjaans{})
    DB.AutoMigrate(&models.User{})
    DB.AutoMigrate(&models.Wargas{})


}

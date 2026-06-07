package user_test

import (
	"log"
	"testing"

	"github.com/Qingruiliu0311/vlog_ddd/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestMigrate(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/vlog?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&user.User{})
}

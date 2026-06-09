package impl_test

import (
	"context"
	"testing"

	"github.com/Qingruiliu0311/vlog_ddd/user"
	"github.com/Qingruiliu0311/vlog_ddd/user/impl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var svc *impl.UserServiceImplement

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/vlog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	svc = &impl.UserServiceImplement{Db: db}
}

func TestRegistry(t *testing.T) {
	u, err := svc.Registry(context.Background(), &user.RegistryReq{
		Email:    "test2@test.com",
		Password: "12345678",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

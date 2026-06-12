package impl_test

import (
	"context"
	"testing"

	"github.com/Qingruiliu0311/vlog_ddd/token"
	"github.com/Qingruiliu0311/vlog_ddd/token/impl"
	userImpl "github.com/Qingruiliu0311/vlog_ddd/user/impl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/vlog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	svc := &impl.TokenServiceImplement{
		Db:   db,
		User: &userImpl.UserServiceImplement{Db: db},
	}
	token.Register(svc)
}

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenReq()
	req.Email = "test3@test.com"
	req.Password = "1234"
	token.GetService().IssueToken(context.Background(), req)
}

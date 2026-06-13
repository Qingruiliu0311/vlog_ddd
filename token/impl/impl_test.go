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
		Db:      db,
		UserSvc: &userImpl.UserServiceImplement{Db: db},
	}
	token.RegisterService(svc)
}

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenReq()
	req.Email = "test3@test.com"
	req.Password = "1234"
	token.GetService().IssueToken(context.Background(), req)
}

func TestRevokeToken(t *testing.T) {
	req := token.NewRevokeTokenReq()
	req.AccessToken = "07e685a5-3d5f-4374-9370-80480d71cf84"
	req.RefreshToken = "ddf4bf2a-037e-4a96-a79b-2f609825e3b0"
	_, err := token.GetService().RevokeToken(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenReq()
	req.AccessToken = "98f0fe5a-6acd-491a-bda0-6389d5d5ffbb"
	tk, err := token.GetService().ValidateToken(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

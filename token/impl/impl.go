package impl

import (
	"context"
	"fmt"

	"github.com/Qingruiliu0311/vlog_ddd/exception"
	"github.com/Qingruiliu0311/vlog_ddd/token"
	"github.com/Qingruiliu0311/vlog_ddd/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var TokenService token.Service = &TokenServiceImplement{}

type TokenServiceImplement struct {
	Db      *gorm.DB
	UserSvc user.AdminUserService
}

func (t *TokenServiceImplement) Init() {
	t.UserSvc = user.GetService()
}

// IssueToken implements [token.Service].
func (t *TokenServiceImplement) IssueToken(ctx context.Context, in *token.IssueTokenReq) (*token.Token, error) {
	existingUser := user.NewDescribeUserReq(in.Email)
	u, err := t.UserSvc.DescribeUser(context.Background(), existingUser)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(in.Password))
	if err != nil {
		return nil, err
	}
	token := token.NewToken(fmt.Sprintf("%d", u.ID)).SetRefEmail(u.Email)
	err = t.Db.Create(token).Error
	if err != nil {
		return nil, err
	}
	return token, nil
}

// RevokeToken implements [token.Service].
func (t *TokenServiceImplement) RevokeToken(ctx context.Context, in *token.RevokeTokenReq) (*token.Token, error) {
	result := t.Db.WithContext(ctx).Where("access_token=? OR refresh_token=?", in.AccessToken, in.RefreshToken).Delete(&token.Token{})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, exception.NewNotFoundRequest("Token not exists")
	}
	return nil, nil
}

// ValidateToken implements [token.Service].
func (t *TokenServiceImplement) ValidateToken(ctx context.Context, in *token.ValidateTokenReq) (*token.Token, error) {
	tk := &token.Token{}
	err := t.Db.WithContext(ctx).Model(&token.Token{}).Where("access_token=?", in.AccessToken).Take(tk).Error
	if err != nil {
		return nil, err
	}
	if token.IsTokenExpired(tk.AccessTokenExpiredAt) {
		return nil, exception.NewBadRequest("Token is expired")
	}

	return tk, nil
}

package impl

import (
	"context"
	"fmt"

	"github.com/Qingruiliu0311/vlog_ddd/token"
	"github.com/Qingruiliu0311/vlog_ddd/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var TokenService token.Service = &TokenServiceImplement{}

type TokenServiceImplement struct {
	Db   *gorm.DB
	User user.AdminUserService
}

func (t *TokenServiceImplement) Init() {
	t.User = user.GetService()
}

// IssueToken implements [token.Service].
func (t *TokenServiceImplement) IssueToken(ctx context.Context, in *token.IssueTokenReq) (*token.Token, error) {
	existingUser := user.NewDescribeUserReq(in.Email)
	u, err := t.User.DescribeUser(context.Background(), existingUser)
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
func (t *TokenServiceImplement) RevokeToken(context.Context, *token.RevokeTokenReq) (*token.Token, error) {
	panic("unimplemented")
}

// ValidateToken implements [token.Service].
func (t *TokenServiceImplement) ValidateToken(context.Context, *token.ValidateTokenReq) (*token.Token, error) {
	panic("unimplemented")
}

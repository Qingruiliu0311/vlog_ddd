package impl

import (
	"context"

	"github.com/Qingruiliu0311/vlog_ddd/token"
)

var TokenService token.Service = &TokenServiceImplement{}

type TokenServiceImplement struct{}

// IssueToken implements [token.Service].
func (t *TokenServiceImplement) IssueToken(context.Context, *token.IssueTokenReq) (*token.Token, error) {
	panic("unimplemented")
}

// RevokeToken implements [token.Service].
func (t *TokenServiceImplement) RevokeToken(context.Context, *token.RevokeTokenReq) (*token.Token, error) {
	panic("unimplemented")
}

// ValidateToken implements [token.Service].
func (t *TokenServiceImplement) ValidateToken(context.Context, *token.ValidateTokenReq) (*token.Token, error) {
	panic("unimplemented")
}

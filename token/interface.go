package token

import "context"

type Service interface {
	InnerService
	UserService
}

type InnerService interface {
	ValidateToken(context.Context, *ValidateTokenReq) (*Token, error)
}

type UserService interface {
	IssueToken(context.Context, *IssueTokenReq) (*Token, error)
	RevokeToken(context.Context, *RevokeTokenReq) (*Token, error)
}

type ValidateTokenReq struct {
	AccessToken string `json:"access_token"`
}
type IssueTokenReq struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}
type RevokeTokenReq struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

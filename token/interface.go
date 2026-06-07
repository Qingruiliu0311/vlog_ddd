package token

import "context"

type Service interface {
	InnerService
	UserService
}

type InnerService interface {
	ValidateToken(context.Context, *ValidateTokenReq)
}

type UserService interface {
	IssueToken(context.Context, *IssueTokenReq)
	RevokeToken(context.Context, *RevokeTokenReq)
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

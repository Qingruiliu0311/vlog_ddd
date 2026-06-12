package token

import (
	"context"
	"time"
)

type Service interface {
	InnerService
	UserService
}

var service Service

func GetService() Service {
	return service
}

func Register(svc Service) {
	service = svc
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

func NewValidateTokenReq() *ValidateTokenReq {
	return &ValidateTokenReq{}
}

type IssueTokenReq struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

func NewIssueTokenReq() *IssueTokenReq {
	return &IssueTokenReq{}
}

type RevokeTokenReq struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewRevokeTokenReq() *RevokeTokenReq {
	return &RevokeTokenReq{}
}

func IsTokenExpired(expiredAt *time.Time) bool {
	now := time.Now()
	if now.After(*expiredAt) {
		return false
	}
	return true
}

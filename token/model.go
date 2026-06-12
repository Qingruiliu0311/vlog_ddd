package token

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	Id        uint   `json:"id" gorm:"column:id;primaryKey"`
	RefUserId string `json:"ref_user_id" gorm:"column:ref_user_id"`
	RefEmail  string `json:"ref_email" gorm:"column:ref_email;uniqueIndex"`

	//access token
	AccessToken          string     `json:"access_token" gorm:"column:access_token;unique;index"`
	IssuedAt             *time.Time `json:"issued_at" gorm:"column:issued_at"`
	AccessTokenExpiredAt *time.Time `json:"access_token_expired_at" gorm:"column:access_token_expired_at"`

	//refresh token
	RefreshToken          string     `json:"refresh_token" gorm:"column:refresh_token;unique;index"`
	RefreshTokenExpiredAt *time.Time `json:"refresh_token_expired_at" gorm:"column:refresh_token_expired_at"`
}

func NewToken(refUserId string) *Token {
	aExpiredAt := time.Now().AddDate(0, 0, 1)
	rExpiredAt := time.Now().AddDate(0, 0, 7)
	return &Token{
		RefUserId:             refUserId,
		AccessToken:           uuid.NewString(),
		AccessTokenExpiredAt:  &aExpiredAt,
		RefreshToken:          uuid.NewString(),
		RefreshTokenExpiredAt: &rExpiredAt,
	}
}

func (t *Token) SetRefEmail(refEmail string) *Token {
	t.RefEmail = refEmail
	return t
}

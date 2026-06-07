package token

import "time"

type Token struct {
	Id          uint   `json:"id" gorm:"column:id;primaryKey"`
	RefUserId   string `json:"ref_user_id" gorm:"column:ref_user_id"`
	RefUserName string `json:"ref_user_name" gorm:"-"`

	//access token
	AccessToken          string     `json:"access_token" gorm:"column:access_token;unique;index"`
	IssuedAt             *time.Time `json:"issued_at" gorm:"column:issued_at"`
	AccessTokenExpiredAt *time.Time `json:"access_token_expired_at" gorm:"column:access_token_expired_at"`

	//refresh token
	RefreshToken          string     `json:"refresh_token" gorm:"column:refresh_token;unique;index"`
	RefreshTokenExpiredAt *time.Time `json:"refresh_token_expired_at" gorm:"column:refresh_token_expired_at"`
}

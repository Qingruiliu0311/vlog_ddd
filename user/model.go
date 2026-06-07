package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Name string `json:"name" gorm:"column:name;type:varchar(255)"`
	Age  uint   `json:"age" gorm:"column:age;type:uint"`
	RegistryReq
	gorm.Model
}

type RegistryReq struct {
	Email    string `json:"email" gorm:"column:username;unique;index" validate:"required"`
	Password string `json:"password" gorm:"column:password;type:varchar(255)" validate:"required"`
	Profile
	Status
}

type Status struct {
	BlockedAt   *time.Time `json:"blocked_at" gorm:"column:blocked_at"`
	BlockReason string     `json:"block_reason" gorm:"column:block_reason;type:text"`
}

func (c *Status) IsBlocked() bool {
	return c.BlockedAt != nil
}

type Profile struct {
	Avatar string `json:"avatar" gorm:"column:avatar;type:varchar(255)"`
}

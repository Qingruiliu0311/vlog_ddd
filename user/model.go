package user

import (
	"time"

	"github.com/Qingruiliu0311/vlog_ddd/exception"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

//TODO: Audit log for block/unblock of a user

func New(req *RegistryReq) (*User, error) {
	err := req.Validate()
	if err != nil {
		return nil, exception.NewBadRequest("参数校验失败: %s", err)
	}
	return &User{
		RegistryReq: *req,
	}, nil
}

type User struct {
	RegistryReq
	gorm.Model
	Status
}

type RegistryReq struct {
	Email    string `json:"email" gorm:"column:email;unique;index" validate:"required"`
	Password string `json:"-" gorm:"column:password;type:varchar(255)" validate:"required"`
	Name     string `json:"name" gorm:"column:name;type:varchar(255)"`
	Age      uint   `json:"age" gorm:"column:age;type:uint"`
	Profile
}

var validate = validator.New()

func (r RegistryReq) Validate() error {
	return validate.Struct(r)
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

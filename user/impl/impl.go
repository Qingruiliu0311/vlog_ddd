package impl

import (
	"context"

	"github.com/Qingruiliu0311/vlog_ddd/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var UserService user.Service = &UserServiceImplement{}

type UserServiceImplement struct {
	Db *gorm.DB
}

// Registry implements [user.Service].
func (u *UserServiceImplement) Registry(ctx context.Context, in *user.RegistryReq) (*user.User, error) {
	ins, err := user.New(in)
	if err != nil {
		return nil, err
	}
	HashPass, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	ins.Password = string(HashPass)
	err = u.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(ins).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// ResetPassword implements [user.Service].
func (u *UserServiceImplement) ResetPassword(context.Context, user.ResetPasswordReq) (*user.User, error) {
	panic("unimplemented")
}

// Unregistry implements [user.Service].
func (u *UserServiceImplement) Unregistry(context.Context, user.UnregistryReq) error {
	panic("unimplemented")
}

// UpdatePassword implements [user.Service].
func (u *UserServiceImplement) UpdatePassword(context.Context, user.UpdatePasswordReq) error {
	panic("unimplemented")
}

// UpdateProfile implements [user.Service].
func (u *UserServiceImplement) UpdateProfile(context.Context) {
	panic("unimplemented")
}

// UpdateUserStatus implements [user.Service].
func (u *UserServiceImplement) UpdateUserStatus(context.Context, user.UpdateUserStatusReq) (*user.User, error) {
	panic("unimplemented")
}

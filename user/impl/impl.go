package impl

import (
	"context"
	"errors"

	"github.com/Qingruiliu0311/vlog_ddd/exception"
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
		var existingUser user.User
		err := tx.Where("email=?", in.Email).First(&existingUser).Error
		if err == nil {
			return exception.NewConflictRequest("Email %s has already being registered", existingUser.Email)
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.NewBadRequest("bad request: %s", err)
		}

		err = tx.Create(ins).Error
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
func (u *UserServiceImplement) ResetPassword(ctx context.Context, ins *user.ResetPasswordReq) (*user.User, error) {
	hashpass, err := bcrypt.GenerateFromPassword([]byte(ins.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	//fake it do it later
	sentcode := "123456"
	if ins.VerificationCode != sentcode {
		return nil, exception.NewUnauthorisation("Incorrect verification code")
	}
	ins.NewPassword = string(hashpass)
	err = u.Db.Model(&user.User{}).Where("email=?", ins.Email).Update("password", ins.NewPassword).Error
	if err != nil {
		return nil, err
	}

	var result = &user.User{}

	err = u.Db.Where("email=?", ins.Email).Take(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Unregistry implements [user.Service].
func (u *UserServiceImplement) Unregistry(ctx context.Context, in *user.UnregistryReq) error {
	err := u.Db.Model(&user.User{}).Where("email=?", in.Email).Delete(&user.User{}).Error
	if err != nil {
		return exception.NewBadRequest("couldn't delete %s", err)
	}
	return nil
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

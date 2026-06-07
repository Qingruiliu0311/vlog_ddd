package impl

import (
	"context"

	"github.com/Qingruiliu0311/vlog_ddd/user"
)

var UserService user.Service = &UserServiceImplement{}

type UserServiceImplement struct{}

// Registry implements [user.Service].
func (u *UserServiceImplement) Registry(context.Context, *user.RegistryReq) (*user.User, error) {
	panic("unimplemented")
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

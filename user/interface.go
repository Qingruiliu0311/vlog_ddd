package user

import "context"

type Service interface {
	AdminUserService
	UserService
}

type AdminUserService interface {
	UpdateUserStatus(context.Context, UpdateUserStatusReq) (*User, error)
}

type UserService interface {
	Registry(context.Context, *RegistryReq) (*User, error)
	UpdatePassword(context.Context, UpdatePasswordReq) error
	ResetPassword(context.Context, *ResetPasswordReq) (*User, error)
	Unregistry(context.Context, UnregistryReq) error
	UpdateProfile(context.Context)
}
type UpdateUserStatusReq struct {
	Email string `json:"email"`
	Status
}

type UpdatePasswordReq struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ResetPasswordReq struct {
	Email            string `json:"email"`
	NewPassword      string `json:"new_password"`
	VerificationCode string `json:"verification_code"`
}

type UnregistryReq struct {
	Email            string `json:"email"`
	VerificationCode string `json:"verification_code"`
}

type UpdateProfileReq struct {
	UserId uint `json:"user_id"`
	Profile
}

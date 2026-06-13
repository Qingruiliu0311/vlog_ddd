package user

import "context"

type Service interface {
	AdminUserService
	UserService
}

var service Service

func GetService() Service {
	return service
}

func RegisterService(s Service) {
	service = s
}

type AdminUserService interface {
	// UpdateUserStatus(context.Context, *UpdateUserStatusReq) (*User, error)
	UpdateUserStatus
}
type UpdateUserStatus interface {
	DescribeUser(context.Context, *DescribeUserReq) (*User, error)
	BlockUserStatus(context.Context, *BlockUserStatusReq) (*User, error)
	UnblockUserStatus(context.Context, *UnblockUserStatusReq) (*User, error)
}
type UserService interface {
	Registry(context.Context, *RegistryReq) (*User, error)
	UpdatePassword(context.Context, *UpdatePasswordReq) (*User, error)
	ResetPassword(context.Context, *ResetPasswordReq) (*User, error)
	Unregistry(context.Context, *UnregistryReq) error
	UpdateProfile(context.Context)
}

type DescribeUserReq struct {
	Email string `json:"email"`
}

func NewDescribeUserReq(email string) *DescribeUserReq {
	return &DescribeUserReq{Email: email}
}

type BlockUserStatusReq struct {
	Email string `json:"email"`
	Status
}

type UnblockUserStatusReq struct {
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

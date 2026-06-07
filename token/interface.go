package token

type Service interface {
	InnerService
	UserService
}

type InnerService interface {
	ValidateToken()
}

type UserService interface {
	IssueToken()
	RevokeToken()
}

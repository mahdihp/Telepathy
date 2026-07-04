package domain

import "context"

type Service interface {
	Register(ctx context.Context, req RegisterRequest) (*AuthResult, error)
	Login(ctx context.Context, req LoginRequest) (*AuthResult, error)
}

type AuthResult struct {
	UserID      string
	AccessToken string
}
type RegisterRequest struct {
	PhoneNumber string
	DisplayName string
}

type LoginRequest struct {
	PhoneNumber string
}

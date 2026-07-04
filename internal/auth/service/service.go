package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	authdomain "github.com/mahdihp/telepathy/internal/auth/domain"
	userdomain "github.com/mahdihp/telepathy/internal/user/domain"
)

import "errors"

var (
	ErrPhoneAlreadyExists = errors.New("phone number already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Service struct {
	users userdomain.UserRepository
}

func (s *Service) Register(ctx context.Context, req authdomain.RegisterRequest) (*authdomain.AuthResult, error) {
	user, err := s.users.FindByPhone(ctx, req.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, ErrPhoneAlreadyExists
	}

	now := time.Now()

	user = &userdomain.User{
		ID:          uuid.NewString(),
		PhoneNumber: req.PhoneNumber,
		DisplayName: req.DisplayName,
		Status:      userdomain.UserStatusOffline,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.users.Create(ctx, user); err != nil {
		return nil, err
	}

	token := "TODO"

	return &authdomain.AuthResult{
		UserID:      user.ID,
		AccessToken: token,
	}, nil
}

func (s *Service) Login(ctx context.Context, req authdomain.LoginRequest) (*authdomain.AuthResult, error) {

	user, err := s.users.FindByPhone(ctx, req.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrInvalidCredentials
	}

	token := "TODO"

	return &authdomain.AuthResult{
		UserID:      user.ID,
		AccessToken: token,
	}, nil
}

func New(users userdomain.UserRepository) authdomain.Service {
	return &Service{
		users: users,
	}
}

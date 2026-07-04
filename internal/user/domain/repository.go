package domain

import "context"

// UserRepository defines the contract for user persistence.
type UserRepository interface {
	Create(ctx context.Context, user *User) error

	FindByID(ctx context.Context, id string) (*User, error)

	FindByPhone(ctx context.Context, phone string) (*User, error)

	Update(ctx context.Context, user *User) error
}

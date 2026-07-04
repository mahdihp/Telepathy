package repository

import "time"

type userModel struct {
	ID          string
	PhoneNumber string
	Username    string
	DisplayName string
	AvatarURL   string
	Bio         string

	Status     string
	LastSeenAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

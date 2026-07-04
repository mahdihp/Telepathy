package domain

import "time"

type UserStatus string

const (
	UserStatusOnline  UserStatus = "online"
	UserStatusOffline UserStatus = "offline"
	UserStatusBusy    UserStatus = "busy"
)

// User represents a user in the system.
type User struct {
	ID          string
	PhoneNumber string
	Username    string
	DisplayName string
	AvatarURL   string
	Bio         string

	Status     UserStatus
	LastSeenAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

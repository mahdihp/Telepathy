package repository

import "github.com/mahdihp/telepathy/internal/user/domain"

func toDomain(m *userModel) *domain.User {
	if m == nil {
		return nil
	}

	return &domain.User{
		ID:          m.ID,
		PhoneNumber: m.PhoneNumber,
		Username:    m.Username,
		DisplayName: m.DisplayName,
		AvatarURL:   m.AvatarURL,
		Bio:         m.Bio,
		Status:      domain.UserStatus(m.Status),
		LastSeenAt:  m.LastSeenAt,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func toModel(u *domain.User) *userModel {
	return &userModel{
		ID:          u.ID,
		PhoneNumber: u.PhoneNumber,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		AvatarURL:   u.AvatarURL,
		Bio:         u.Bio,
		Status:      string(u.Status),
		LastSeenAt:  u.LastSeenAt,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

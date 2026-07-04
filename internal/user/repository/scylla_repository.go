package repository

import (
	"context"

	gocql "github.com/apache/cassandra-gocql-driver/v2"

	"github.com/mahdihp/telepathy/internal/user/domain"
)

type ScyllaRepository struct {
	session *gocql.Session
}

func NewScyllaRepository(session *gocql.Session) domain.UserRepository {
	return &ScyllaRepository{
		session: session,
	}
}

func (r *ScyllaRepository) Create(ctx context.Context, user *domain.User) error {
	//TODO implement me
	query := `INSERT INTO users_by_id (
		user_id,
		phone_number,
		username,
		display_name,
		avatar_url,
		bio,
		status,
		last_seen_at,
		created_at,
		updated_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	if err := r.session.Query(
		query,
		user.ID,
		user.PhoneNumber,
		user.Username,
		user.DisplayName,
		user.AvatarURL,
		user.Bio,
		string(user.Status),
		user.LastSeenAt,
		user.CreatedAt,
		user.UpdatedAt,
	).Exec(); err != nil {
		return err
	}

	query = `
	INSERT INTO users_by_phone (
		phone_number,
		user_id
	) VALUES (?, ?)`

	return r.session.Query(
		query,
		user.PhoneNumber,
		user.ID,
	).Exec()

}

func (r *ScyllaRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	query := `
	SELECT
		user_id,
		phone_number,
		username,
		display_name,
		avatar_url,
		bio,
		status,
		last_seen_at,
		created_at,
		updated_at
	FROM users_by_id
	WHERE user_id = ?
	LIMIT 1`

	var user domain.User
	var status string

	err := r.session.Query(query, id).Scan(
		&user.ID,
		&user.PhoneNumber,
		&user.Username,
		&user.DisplayName,
		&user.AvatarURL,
		&user.Bio,
		&status,
		&user.LastSeenAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}

	user.Status = domain.UserStatus(status)

	return &user, nil
}

func (r *ScyllaRepository) FindByPhone(ctx context.Context, phone string) (*domain.User, error) {
	query := `
	SELECT user_id
	FROM users_by_phone
	WHERE phone_number = ?
	LIMIT 1`

	var userID string

	err := r.session.Query(query, phone).Scan(&userID)

	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}

	return r.FindByID(ctx, userID)
}

func (r *ScyllaRepository) Update(ctx context.Context, user *domain.User) error {
	query := `
	UPDATE users_by_id
	SET
		username = ?,
		display_name = ?,
		avatar_url = ?,
		bio = ?,
		status = ?,
		last_seen_at = ?,
		updated_at = ?
	WHERE user_id = ?`

	return r.session.Query(
		query,
		user.Username,
		user.DisplayName,
		user.AvatarURL,
		user.Bio,
		string(user.Status),
		user.LastSeenAt,
		user.UpdatedAt,
		user.ID,
	).Exec()
}

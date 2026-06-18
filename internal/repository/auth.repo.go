package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aqilknz/shortlink-backend/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type AuthRepository interface {
	CheckEmailExists(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	IsTokenBlacklisted(ctx context.Context, userID int, token string) bool
	AddTokenToBlacklist(ctx context.Context, userID int, token string, expiresIn time.Duration) error
}

type authRepository struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

func NewAuthRepository(db *pgxpool.Pool, redis *redis.Client) AuthRepository {
	return &authRepository{
		db:    db,
		redis: redis,
	}
}

func (ar *authRepository) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	sql := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	var exists bool
	err := ar.db.QueryRow(ctx, sql, email).Scan(&exists)
	return exists, err
}

func (ar *authRepository) CreateUser(ctx context.Context, user *model.User) error {
	sql := `
		INSERT INTO users (email, password_hash) 
		VALUES ($1, $2) 
		RETURNING id, created_at
	`
	args := []any{user.Email, user.Password}

	err := ar.db.QueryRow(ctx, sql, args...).Scan(
		&user.Id,
		&user.CreatedAt,
	)

	return err
}

func (ar *authRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	sql := `SELECT id, email, password_hash, full_name, avatar_url, created_at FROM users WHERE email = $1`
	args := []any{email}

	var user model.User
	err := ar.db.QueryRow(ctx, sql, args...).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.ProfilePhoto,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

func (ar *authRepository) AddTokenToBlacklist(ctx context.Context, userID int, token string, expiresIn time.Duration) error {
	key := fmt.Sprintf("blacklist:user:%d:token:%s", userID, token)
	log.Printf("DEBUG: : %s, TTL: %v", key, expiresIn)
	return ar.redis.Set(ctx, key, "revoked", expiresIn).Err()
}

func (ar *authRepository) IsTokenBlacklisted(ctx context.Context, userID int, token string) bool {
	key := fmt.Sprintf("blacklist:user:%d:token:%s", userID, token)
	err := ar.redis.Get(ctx, key).Err()

	if errors.Is(err, redis.Nil) {
		return false
	}

	return true
}

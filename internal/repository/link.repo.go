package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/aqilknz/shortlink-backend/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LinkRepository interface {
	CreateLink(ctx context.Context, link *model.Link) error
	GetLinksByUserID(ctx context.Context, userID int, search string, limit int, offset int) ([]model.Link, error)
	CountLinksByUserID(ctx context.Context, userID int, search string) (int, error)
	GetLinkBySlug(ctx context.Context, slug string) (model.Link, error)
	DeleteLink(ctx context.Context, linkID int, userID int) error
	IncrementClickCount(ctx context.Context, slug string) error
}

type linkRepository struct {
	db *pgxpool.Pool
}

func NewLinkRepository(db *pgxpool.Pool) LinkRepository {
	return &linkRepository{db: db}
}

func (r *linkRepository) CreateLink(ctx context.Context, link *model.Link) error {
	sql := `
		INSERT INTO links (user_id, original_url, slug) 
		VALUES ($1, $2, $3) 
		RETURNING id, created_at
	`
	err := r.db.QueryRow(ctx, sql, link.UserId, link.OriginalURL, link.Slug).Scan(
		&link.Id,
		&link.CreatedAt,
	)
	return err
}

func (r *linkRepository) GetLinksByUserID(ctx context.Context, userID int, search string, limit int, offset int) ([]model.Link, error) {
	query := `
		SELECT id, user_id, original_url, slug, click_count, created_at 
		FROM links 
		WHERE user_id = $1 AND deleted_at IS NULL
	`
	args := []interface{}{userID}
	argID := 2

	if search != "" {
		query += fmt.Sprintf(` AND (original_url ILIKE $%d OR slug ILIKE $%d)`, argID, argID)
		args = append(args, "%"+search+"%")
		argID++
	}

	query += fmt.Sprintf(` ORDER BY created_at DESC LIMIT $%d OFFSET $%d`, argID, argID+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []model.Link
	for rows.Next() {
		var link model.Link
		err := rows.Scan(&link.Id, &link.UserId, &link.OriginalURL, &link.Slug, &link.ClickCount, &link.CreatedAt)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}
	return links, nil
}

func (r *linkRepository) CountLinksByUserID(ctx context.Context, userID int, search string) (int, error) {
	query := `SELECT COUNT(id) FROM links WHERE user_id = $1 AND deleted_at IS NULL`
	args := []interface{}{userID}

	if search != "" {
		query += ` AND (original_url ILIKE $2 OR slug ILIKE $2)`
		args = append(args, "%"+search+"%")
	}

	var total int
	err := r.db.QueryRow(ctx, query, args...).Scan(&total)
	return total, err
}

func (r *linkRepository) GetLinkBySlug(ctx context.Context, slug string) (model.Link, error) {
	sql := `
		SELECT id, user_id, original_url, slug, created_at 
		FROM links 
		WHERE slug = $1 AND deleted_at IS NULL
	`
	var link model.Link
	err := r.db.QueryRow(ctx, sql, slug).Scan(
		&link.Id, &link.UserId, &link.OriginalURL, &link.Slug, &link.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return link, errors.New("link not found")
		}
		return link, err
	}
	return link, nil
}

func (r *linkRepository) DeleteLink(ctx context.Context, linkID int, userID int) error {
	sql := `
		UPDATE links 
		SET deleted_at = CURRENT_TIMESTAMP 
		WHERE id = $1 AND user_id = $2 AND deleted_at IS NULL
	`
	commandTag, err := r.db.Exec(ctx, sql, linkID, userID)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("link not found or unauthorized")
	}
	return nil
}

func (r *linkRepository) IncrementClickCount(ctx context.Context, slug string) error {
	query := `UPDATE links SET click_count = click_count + 1 WHERE slug = $1`
	_, err := r.db.Exec(ctx, query, slug)
	return err
}

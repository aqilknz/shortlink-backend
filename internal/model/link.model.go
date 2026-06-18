package model

import "time"

type Link struct {
	Id          int        `json:"id"`
	UserId      int        `json:"user_id"`
	OriginalURL string     `json:"original_url"`
	Slug        string     `json:"slug"`
	ClickCount  int        `db:"click_count"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

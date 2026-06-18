package dto

import "time"

type CreateLinkRequest struct {
	OriginalURL string `json:"original_url" binding:"required,url" example:"https://www.linkedin.com/in/ahmadaqil"`
	CustomSlug  string `json:"custom_slug,omitempty" example:"aqil-profil"`
}

type LinkResponse struct {
	Id          int       `json:"id" example:"1"`
	OriginalURL string    `json:"original_url" example:"https://www.linkedin.com/in/ahmadaqil"`
	Slug        string    `json:"slug" example:"aqil-profil"`
	ClickCount  int       `json:"click_count" example:"42"`
	CreatedAt   time.Time `json:"created_at" example:"2026-06-17T15:04:05Z"`
}

type PaginatedLinkResponse struct {
	Data []LinkResponse `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	CurrentPage  int `json:"current_page" example:"1"`
	Limit        int `json:"limit" example:"10"`
	TotalRecords int `json:"total_records" example:"45"`
	TotalPages   int `json:"total_pages" example:"5"`
}

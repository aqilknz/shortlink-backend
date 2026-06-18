package service

import (
	"context"
	"regexp"
	"strings"

	"github.com/aqilknz/shortlink-backend/internal/dto"
	"github.com/aqilknz/shortlink-backend/internal/model"
	"github.com/aqilknz/shortlink-backend/internal/repository"
	"github.com/aqilknz/shortlink-backend/internal/utils"
)

type LinkService interface {
	CreateLink(ctx context.Context, userID int, req dto.CreateLinkRequest) (dto.LinkResponse, error)
	GetUserLinks(ctx context.Context, userID int, search string, page int, limit int) (dto.PaginatedLinkResponse, error)
	DeleteLink(ctx context.Context, userID int, linkID int) error
	GetOriginalURL(ctx context.Context, slug string) (string, error)
	CheckSlugAvailability(ctx context.Context, slug string) (bool, error)
}

type linkService struct {
	linkRepo repository.LinkRepository
}

func NewLinkService(linkRepo repository.LinkRepository) LinkService {
	return &linkService{linkRepo: linkRepo}
}

func (ls *linkService) CreateLink(ctx context.Context, userID int, req dto.CreateLinkRequest) (dto.LinkResponse, error) {
	var finalSlug string

	if req.CustomSlug != "" {
		if len(req.CustomSlug) < 3 || len(req.CustomSlug) > 50 {
			return dto.LinkResponse{}, utils.ErrInvalidSlugLength
		}

		if !regexp.MustCompile(`^[a-zA-Z0-9\-]+$`).MatchString(req.CustomSlug) {
			return dto.LinkResponse{}, utils.ErrInvalidSlugFormat
		}

		switch strings.ToLower(req.CustomSlug) {
		case "api", "login", "register", "dashboard":
			return dto.LinkResponse{}, utils.ErrReservedSlug
		}

		if _, err := ls.linkRepo.GetLinkBySlug(ctx, req.CustomSlug); err == nil {
			return dto.LinkResponse{}, utils.ErrSlugAlreadyExists
		}

		finalSlug = req.CustomSlug
	} else {
		finalSlug = utils.GenerateSlug(6)
	}

	link := model.Link{
		UserId:      userID,
		OriginalURL: req.OriginalURL,
		Slug:        finalSlug,
	}

	if err := ls.linkRepo.CreateLink(ctx, &link); err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "23505") {
			return dto.LinkResponse{}, utils.ErrSlugAlreadyExists
		}
		return dto.LinkResponse{}, utils.ErrInternalServer
	}

	return dto.LinkResponse{
		Id:          link.Id,
		OriginalURL: link.OriginalURL,
		Slug:        link.Slug,
		ClickCount:  0, // Default saat baru dibuat
		CreatedAt:   link.CreatedAt,
	}, nil
}

// UPDATE: Implementasi Paginasi dan Search
func (ls *linkService) GetUserLinks(ctx context.Context, userID int, search string, page int, limit int) (dto.PaginatedLinkResponse, error) {
	// 1. Hitung Offset
	offset := (page - 1) * limit

	// 2. Ambil total data untuk metadata paginasi
	totalRecords, err := ls.linkRepo.CountLinksByUserID(ctx, userID, search)
	if err != nil {
		return dto.PaginatedLinkResponse{}, utils.ErrInternalServer
	}

	// 3. Ambil data link
	links, err := ls.linkRepo.GetLinksByUserID(ctx, userID, search, limit, offset)
	if err != nil {
		return dto.PaginatedLinkResponse{}, utils.ErrInternalServer
	}

	// 4. Hitung total halaman
	totalPages := totalRecords / limit
	if totalRecords%limit != 0 {
		totalPages++
	}

	// 5. Mapping ke DTO
	var resData []dto.LinkResponse
	for _, l := range links {
		resData = append(resData, dto.LinkResponse{
			Id:          l.Id,
			OriginalURL: l.OriginalURL,
			Slug:        l.Slug,
			ClickCount:  l.ClickCount,
			CreatedAt:   l.CreatedAt,
		})
	}

	if resData == nil {
		resData = []dto.LinkResponse{}
	}

	return dto.PaginatedLinkResponse{
		Data: resData,
		Meta: dto.PaginationMeta{
			CurrentPage:  page,
			Limit:        limit,
			TotalRecords: totalRecords,
			TotalPages:   totalPages,
		},
	}, nil
}

func (ls *linkService) DeleteLink(ctx context.Context, userID int, linkID int) error {
	err := ls.linkRepo.DeleteLink(ctx, linkID, userID)
	if err != nil {
		if err.Error() == "link not found or unauthorized" {
			return utils.ErrForbidden
		}
		return utils.ErrInternalServer
	}
	return nil
}

func (ls *linkService) GetOriginalURL(ctx context.Context, slug string) (string, error) {
	link, err := ls.linkRepo.GetLinkBySlug(ctx, slug)
	if err != nil {
		return "", utils.ErrNotFound
	}

	// UPDATE: Trik background task (Goroutine) untuk menambah jumlah klik
	go func() {
		// Gunakan context baru agar tidak terputus saat request utama selesai
		_ = ls.linkRepo.IncrementClickCount(context.Background(), slug)
	}()

	return link.OriginalURL, nil
}

func (ls *linkService) CheckSlugAvailability(ctx context.Context, slug string) (bool, error) {
	_, err := ls.linkRepo.GetLinkBySlug(ctx, slug)
	if err != nil {
		return true, nil // Tidak error artinya belum dipakai = Tersedia
	}
	return false, nil // Tidak ada error berarti datanya ketemu = Tidak Tersedia
}

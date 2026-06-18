package service

import (
	"context"
	"time"

	"github.com/aqilknz/shortlink-backend/internal/dto"
	"github.com/aqilknz/shortlink-backend/internal/model"
	"github.com/aqilknz/shortlink-backend/internal/repository"
	"github.com/aqilknz/shortlink-backend/internal/utils"
	"github.com/aqilknz/shortlink-backend/pkg"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterRequest) error
	Login(ctx context.Context, req dto.LoginRequest) (dto.AuthResponse, error)
	Logout(ctx context.Context, userID int, token string, expiresAt time.Time) error
}

type authService struct {
	authRepo   repository.AuthRepository
	hashConfig *pkg.HashConfig
}

func NewAuthService(authRepo repository.AuthRepository, hashConfig *pkg.HashConfig) AuthService {
	return &authService{
		authRepo:   authRepo,
		hashConfig: hashConfig,
	}
}

func (s *authService) Register(ctx context.Context, req dto.RegisterRequest) error {
	if !pkg.IsValidEmail(req.Email) {
		return utils.ErrInvalidEmail
	}

	exists, err := s.authRepo.CheckEmailExists(ctx, req.Email)
	if err != nil {
		return utils.ErrInternalServer
	}
	if exists {
		return utils.ErrEmailAlreadyExists
	}

	hashedPassword := s.hashConfig.GenHash(req.Password)

	newUser := &model.User{
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := s.authRepo.CreateUser(ctx, newUser); err != nil {
		return utils.ErrInternalServer
	}

	return nil
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (dto.AuthResponse, error) {
	if !pkg.IsValidEmail(req.Email) {
		return dto.AuthResponse{}, utils.ErrInvalidEmail
	}

	user, err := s.authRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return dto.AuthResponse{}, utils.ErrInvalidCredentials
	}

	if err := s.hashConfig.Compare(req.Password, user.Password); err != nil {
		return dto.AuthResponse{}, utils.ErrInvalidCredentials
	}

	claims := pkg.NewClaims(user.Id, user.FullName)
	token, err := claims.GenJWT()
	if err != nil {
		return dto.AuthResponse{}, utils.ErrInternalServer
	}

	res := dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:        user.Id,
			Email:     user.Email,
			FullName:  user.FullName,
			AvatarURL: user.ProfilePhoto,
		},
	}

	return res, nil
}

func (s *authService) Logout(ctx context.Context, userID int, token string, expiresAt time.Time) error {
	if token == "" {
		return utils.ErrInvalidInput
	}

	duration := time.Until(expiresAt)
	if duration <= 0 {
		return nil
	}

	err := s.authRepo.AddTokenToBlacklist(ctx, userID, token, duration)
	if err != nil {
		return utils.ErrInternalServer
	}

	return nil
}

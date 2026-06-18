package controller

import (
	"net/http"
	"time"

	"github.com/aqilknz/shortlink-backend/internal/dto"
	"github.com/aqilknz/shortlink-backend/internal/service"
	"github.com/aqilknz/shortlink-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Register
// @Summary      Register a new user
// @Description  Registers a new user with an email and password.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.RegisterRequest true "Register Payload. Example: {'email': 'user@example.com', 'password': 'password123'}"
// @Success      201  {object}  utils.BaseResponse "Success example: {'status': 201, 'message': 'Registration successful', 'results': null}"
// @Failure      400  {object}  utils.BaseResponse "Error example: Invalid input data format"
// @Failure      409  {object}  utils.BaseResponse "Error example: Email is already registered"
// @Failure      500  {object}  utils.BaseResponse "Error example: Internal server error"
// @Router       /auth/register [post]
func (ctrl *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, utils.ErrInvalidInput.Code, utils.ErrInvalidInput.Message)
		return
	}

	err := ctrl.authService.Register(c.Request.Context(), req)
	if err != nil {
		if customErr, ok := err.(*utils.AppError); ok {
			utils.ErrorResponse(c, customErr.Code, customErr.Message)
			return
		}
		utils.ErrorResponse(c, utils.ErrInternalServer.Code, utils.ErrInternalServer.Message)
		return
	}

	// Return Success
	utils.SuccessResponse(c, http.StatusCreated, "Registration successful", nil)
}

// Login
// @Summary      Login user
// @Description  Authenticates a user using email and password, returning a JWT token.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginRequest true "Login Payload. Example: {'email': 'user@example.com', 'password': 'password123'}"
// @Success      200  {object}  utils.BaseResponse{results=dto.AuthResponse} "Success example: {'status': 200, 'message': 'Login successful', 'results': {'token': 'eyJhbG...', 'user': {'id': 1, 'email': 'user@example.com'}}}"
// @Failure      400  {object}  utils.BaseResponse "Error example: Invalid input data format"
// @Failure      401  {object}  utils.BaseResponse "Error example: Invalid email or password"
// @Failure      500  {object}  utils.BaseResponse "Error example: Internal server error"
// @Router       /auth/login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, utils.ErrInvalidInput.Code, utils.ErrInvalidInput.Message)
		return
	}

	res, err := ctrl.authService.Login(c.Request.Context(), req)
	if err != nil {
		if customErr, ok := err.(*utils.AppError); ok {
			utils.ErrorResponse(c, customErr.Code, customErr.Message)
			return
		}
		utils.ErrorResponse(c, utils.ErrInternalServer.Code, utils.ErrInternalServer.Message)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Login successful", res)
}

// Logout
// @Summary      Logout user
// @Description  Invalidates the current user session by blacklisting the JWT token in Redis.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {object}  utils.BaseResponse "Success example: {'status': 200, 'message': 'Logout successful', 'results': null}"
// @Failure      401  {object}  utils.BaseResponse "Error example: Invalid or expired session, please login again"
// @Failure      500  {object}  utils.BaseResponse "Error example: Internal server error"
// @Router       /auth/logout [delete]
func (ctrl *AuthController) Logout(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.ErrorResponse(c, utils.ErrUnauthorized.Code, utils.ErrUnauthorized.Message)
		return
	}
	exp, exists := c.Get("exp")
	if !exists {
		utils.ErrorResponse(c, utils.ErrUnauthorized.Code, utils.ErrUnauthorized.Message)
		return
	}
	tokenRaw, exists := c.Get("token")
	if !exists {
		utils.ErrorResponse(c, utils.ErrUnauthorized.Code, utils.ErrUnauthorized.Message)
		return
	}

	uid := userID.(int)
	expiresAt := time.Unix(int64(exp.(float64)), 0)
	tokenString := tokenRaw.(string)

	err := ctrl.authService.Logout(c.Request.Context(), uid, tokenString, expiresAt)
	if err != nil {
		if customErr, ok := err.(*utils.AppError); ok {
			utils.ErrorResponse(c, customErr.Code, customErr.Message)
			return
		}
		utils.ErrorResponse(c, utils.ErrInternalServer.Code, "Failed to process logout")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Logout successful", nil)
}

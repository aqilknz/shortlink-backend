package middleware

import (
	"strings"

	"github.com/aqilknz/shortlink-backend/internal/repository"
	"github.com/aqilknz/shortlink-backend/internal/utils"
	"github.com/aqilknz/shortlink-backend/pkg"
	"github.com/gin-gonic/gin"
)

func RequireAuth(authRepo repository.AuthRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			utils.ErrorResponse(ctx, utils.ErrUnauthorized.Code, "Authorization header is missing or invalid")
			ctx.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := pkg.VerifyJWT(tokenString)
		if err != nil {
			utils.ErrorResponse(ctx, utils.ErrUnauthorized.Code, utils.ErrUnauthorized.Message)
			ctx.Abort()
			return
		}

		if authRepo.IsTokenBlacklisted(ctx.Request.Context(), claims.Id, tokenString) {
			utils.ErrorResponse(ctx, utils.ErrUnauthorized.Code, "Session has expired, please login again")
			ctx.Abort()
			return
		}

		ctx.Set("userID", claims.Id)
		ctx.Set("token", tokenString)

		if claims.ExpiresAt != nil {
			ctx.Set("exp", float64(claims.ExpiresAt.Unix()))
		}
		ctx.Next()
	}
}

package middleware

import "github.com/gin-gonic/gin"

func CORSMiddleware(ctx *gin.Context) {
	// Hanya daftar URL Frontend
	allowedOrigins := []string{"http://localhost:5173", "http://localhost:3000"}

	currentOrigin := ctx.GetHeader("Origin")
	for _, origin := range allowedOrigins {
		if currentOrigin == origin {
			ctx.Header("Access-Control-Allow-Origin", currentOrigin)
			break
		}
	}

	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")

	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")

	// Jika ini adalah preflight request (OPTIONS), langsung kembalikan status 204
	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}
	ctx.Next()
}

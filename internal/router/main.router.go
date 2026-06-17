package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func InitRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	// Endpoint untuk testing (Health Check)
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "Database & Redis Connected!",
		})
	})

}

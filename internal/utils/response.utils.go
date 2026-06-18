package utils

import "github.com/gin-gonic/gin"

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

func SuccessResponse(ctx *gin.Context, statusCode int, message string, results interface{}) {
	ctx.JSON(statusCode, BaseResponse{
		Success: true,
		Message: message,
		Results: results,
	})
}

func ErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, BaseResponse{
		Success: false,
		Message: message,
		Results: nil,
	})
}

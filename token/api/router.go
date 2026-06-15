package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(rg *gin.RouterGroup, h *TokenApiHandler) *gin.RouterGroup {
	tokenRoute := rg.Group("/token")
	tokenRoute.POST("", h.IssueToken)
	tokenRoute.POST("/validate", h.ValidateToken)
	return rg
}

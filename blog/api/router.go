package api

import (
	"github.com/Qingruiliu0311/vlog_ddd/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(rg *gin.RouterGroup, h *BlogApiHandler) *gin.RouterGroup {
	blogRoute := rg.Group("/blog")
	blogRoute.Use(middleware.Authentication)
	blogRoute.POST("", h.CreateBlog)
	blogRoute.GET("", h.QueryBlog)
	return rg
}

package api

import (
	"net/http"

	"github.com/Qingruiliu0311/vlog_ddd/blog"
	"github.com/Qingruiliu0311/vlog_ddd/exception"
	"github.com/gin-gonic/gin"
)

type BlogApiHandler struct {
	blogSvc blog.Service
}

func NewBlogApiHandler() *BlogApiHandler {
	return &BlogApiHandler{blogSvc: blog.GetService()}
}

func (b *BlogApiHandler) CreateBlog(ctx *gin.Context) {
	req := &blog.CreateBlogReq{}
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exception.NewBadRequest("请求体解析失败: %s", err))
		return
	}
	ins, err := b.blogSvc.CreateBlog(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.NewInternalServerRequest("create blog in DB failed %s", err))
		return
	}
	ctx.JSON(http.StatusCreated, ins)
}

func (b *BlogApiHandler) QueryBlog(ctx *gin.Context) {
	req := &blog.QueryBlogReq{}
	err := ctx.ShouldBindQuery(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exception.NewBadRequest("请求体解析失败: %s", err))
		return
	}
	ins, err := b.blogSvc.QueryBlog(ctx.Request.Context(), *req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.NewInternalServerRequest("create blog in DB failed %s", err))
		return
	}
	ctx.JSON(http.StatusOK, ins)
}

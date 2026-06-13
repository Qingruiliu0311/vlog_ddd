package blog

import (
	"context"
	"encoding/json"
)

type Service interface {
	CreateBlog(context.Context, *CreateBlogReq) (*Blog, error)
	EditBlog(context.Context, *EditBlogReq) (*Blog, error)
	PublishBlog(context.Context, *PublishBlogReq) (*Blog, error)
	DeleteBlog(context.Context, DeleteBlogReq) error
	QueryBlog(context.Context, QueryBlogReq) (*BlogSet, error)
	DescribeBlog(context.Context) (*Blog, error)
}

var svc Service

func GetService() Service {
	return svc
}

func RegisterService(s Service) {
	svc = s
}

type BlogSet struct {
	Total int64   `json:"total"`
	Items []*Blog `json:"items"`
}

func (b *BlogSet) String() string {
	bs, _ := json.MarshalIndent(b, "", "  ")
	return string(bs)
}

func NewBlogSet() *BlogSet {
	return &BlogSet{}
}

type EditBlogReq struct {
	Id uint `json:"id"`
	CreateBlogReq
}

type PublishBlogReq struct {
	Id uint `json:"id"`
	StatusSpec
}

type DeleteBlogReq struct {
	Id uint `json:"id"`
}

type QueryBlogReq struct {
	Stage     *STAGE            `json:"stage"`
	Keyword   string            `json:"keywords" form:"keywords"`
	CreatedBy string            `json:"created_by" form:"created_by"`
	Category  string            `json:"category" form:"category"`
	Tag       map[string]string `json:"tag" form:"-"`
	Offset    uint              `json:"offset" form:"offset"`
	PageSize  uint              `json:"page_size" form:"page_size"`
}

type DescribeBlogReq struct {
	Id uint `json:"id"`
}

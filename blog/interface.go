package blog

import "context"

type Service interface {
	CreateBlog(context.Context, *CreateBlogReq) (*Blog, error)
	EditBlog(context.Context, *EditBlogReq) (*Blog, error)
	PublishBlog(context.Context, *PublishBlogReq) (*Blog, error)
	DeleteBlog(context.Context, DeleteBlogReq) error
	QueryBlog(context.Context, QueryBlogReq) (*BlogSet, error)
	DescribeBlog(context.Context) (*Blog, error)
}

type BlogSet struct {
	Total int64   `json:"total"`
	Items []*Blog `json:"items"`
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
}

type DescribeBlogReq struct {
	Id uint `json:"id"`
}

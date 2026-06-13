package impl

import (
	"context"

	"github.com/Qingruiliu0311/vlog_ddd/blog"
	"gorm.io/gorm"
)

var BlogService blog.Service = &BlogServiceImplementation{}

type BlogServiceImplementation struct {
	Db *gorm.DB
}

// CreateBlog implements [blog.Service].
func (b *BlogServiceImplementation) CreateBlog(ctx context.Context, in *blog.CreateBlogReq) (*blog.Blog, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	ins := blog.NewCreateBlog(in)
	err = b.Db.Create(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil

}

// DeleteBlog implements [blog.Service].
func (b *BlogServiceImplementation) DeleteBlog(context.Context, blog.DeleteBlogReq) error {
	panic("unimplemented")
}

// DescribeBlog implements [blog.Service].
func (b *BlogServiceImplementation) DescribeBlog(context.Context) (*blog.Blog, error) {
	panic("unimplemented")
}

// EditBlog implements [blog.Service].
func (b *BlogServiceImplementation) EditBlog(context.Context, *blog.EditBlogReq) (*blog.Blog, error) {
	panic("unimplemented")
}

// PublishBlog implements [blog.Service].
func (b *BlogServiceImplementation) PublishBlog(context.Context, *blog.PublishBlogReq) (*blog.Blog, error) {
	panic("unimplemented")
}

// QueryBlog implements [blog.Service].
func (b *BlogServiceImplementation) QueryBlog(context.Context, blog.QueryBlogReq) (*blog.BlogSet, error) {
	panic("unimplemented")
}

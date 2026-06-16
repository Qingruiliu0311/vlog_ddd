package impl

import (
	"context"
	"fmt"

	"github.com/Qingruiliu0311/vlog_ddd/blog"
	"github.com/Qingruiliu0311/vlog_ddd/token"
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
	err = b.Db.WithContext(ctx).Create(ins).Error
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
func (b *BlogServiceImplementation) QueryBlog(ctx context.Context, in blog.QueryBlogReq) (*blog.BlogSet, error) {
	query := b.Db.WithContext(ctx).Model(&blog.Blog{})
	if in.Category != "" {
		query = query.Where("category=?", in.Category)
	}
	if in.CreatedBy != "" {
		query = query.Where("category=?", in.CreatedBy)
	}
	if in.Keyword != "" {
		query = query.Where("title Like ?", "%"+in.Keyword+"%")
	}
	if in.Stage != nil {
		query = query.Where("stage=?", in.Stage)
	}
	if len(in.Tag) > 0 {
		for k, v := range in.Tag {
			query = query.Where("JSON_EXTRACT(tag,?)=?", "$."+k, v)
		}
	}
	fmt.Println(ctx.Value(token.TokenCtxKey{}))

	set := blog.NewBlogSet()

	//total
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}
	//pagination
	err = query.Order("created_at DESC").Offset(int((in.Offset - 1) * in.PageSize)).Limit(int(in.PageSize)).Find(&set.Items).Error
	if err != nil {
		return nil, err
	}
	return set, nil
}

package blog

import "context"

type Service interface {
	CreateBlog(context.Context)
	EditBlog(context.Context)
	PublishBlog(context.Context)
	DeleteBlog(context.Context)
	QueryBlog(context.Context)
	DescribeBlog(context.Context)
}

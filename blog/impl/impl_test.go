package impl_test

import (
	"context"
	"testing"

	"github.com/Qingruiliu0311/vlog_ddd/blog"
	"github.com/Qingruiliu0311/vlog_ddd/blog/impl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/vlog?charset=utf8mb4&parseTime=True&loc=Local"

	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	svc := impl.BlogServiceImplementation{
		Db: Db,
	}
	blog.RegisterService(&svc)
}

func TestCreateBlog(t *testing.T) {
	req := blog.CreateBlogReq{
		Title:   "Golang",
		Content: "This is golang tutorial",
		Summary: "Very Good",
		Tag:     map[string]string{"language": "Golang", "language2": "Java"},
		Catelog: "static",
	}
	b, err := blog.GetService().CreateBlog(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(b)
}

func TestQueryBlog(t *testing.T) {
	stage := blog.Draft
	req := blog.QueryBlogReq{
		Stage:    &stage,
		Offset:   1,
		PageSize: 2,
	}
	result, err := blog.GetService().QueryBlog(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

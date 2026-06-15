package main

import (
	"github.com/Qingruiliu0311/vlog_ddd/blog"
	blogapi "github.com/Qingruiliu0311/vlog_ddd/blog/api"
	blogimpl "github.com/Qingruiliu0311/vlog_ddd/blog/impl"
	"github.com/Qingruiliu0311/vlog_ddd/token"
	tokenapi "github.com/Qingruiliu0311/vlog_ddd/token/api"
	tokenimpl "github.com/Qingruiliu0311/vlog_ddd/token/impl"
	"github.com/Qingruiliu0311/vlog_ddd/user"
	userimpl "github.com/Qingruiliu0311/vlog_ddd/user/impl"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//initiate db
	dsn := "root:123456@tcp(127.0.0.1:3306)/vlog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//register svc
	blogSvc := blogimpl.BlogServiceImplementation{Db: db}
	blog.RegisterService(&blogSvc)

	userSvc := userimpl.UserServiceImplement{Db: db}
	user.RegisterService(&userSvc)

	tokenSvc := tokenimpl.TokenServiceImplement{
		Db:      db,
		UserSvc: user.GetService(),
	}
	token.RegisterService(&tokenSvc)

	// route
	r := gin.Default()
	v1 := r.Group("/api/v1")
	bh := blogapi.NewBlogApiHandler(blog.GetService())
	blogapi.RegisterRouter(v1, bh)

	th := tokenapi.NewTokenApiHandler(token.GetService())
	tokenapi.RegisterRouter(v1, th)
	// token.POST()
	r.Run(":8080")
}

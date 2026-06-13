package blog

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// user, status, title, content, created at, summary, catelog,tag

type Blog struct {
	gorm.Model
	CreateBlogReq
	Status
}

type CreateBlogReq struct {
	Title   string            `json:"title" gorm:"column:title;type:varchar(100)" validate:"required"`
	Content string            `json:"content" gorm:"column:content;type:text" validate:"required"`
	Summary string            `json:"summary" gorm:"column:summary;type:varchar(255)" validate:"required"`
	Tag     map[string]string `json:"tags" gorm:"column:tag;serializer:json"`
	Catelog string            `json:"catelog" gorm:"column:catelog;type:varchar(100)"`
}

func (c *Blog) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return string(b)
}

var validate = validator.New()

func (req *CreateBlogReq) Validate() error {
	return validate.Struct(req)
}

func NewCreateBlog(in *CreateBlogReq) *Blog {
	return &Blog{
		CreateBlogReq: *in,
	}
}

type Status struct {
	StatusSpec
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type StatusSpec struct {
	Stage STAGE `json:"stage" gorm:"column:stage;type:tinyint(1);index"`
}

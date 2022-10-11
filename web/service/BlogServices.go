package service

import (
	"jiuban/model"
	"jiuban/repo"

	"github.com/bwmarrin/snowflake"
	"github.com/kataras/iris/v12"
)

type BlogSerivce interface {
	CreateBlog(ctx iris.Context, data *model.Blog) (out *model.Blog, err error)
	GetBlog(ctx iris.Context) (out []*model.Blog, err error)
	UpdateBlog(ctx iris.Context) (out []*model.Blog, err error)
	DeleteBlog(ctx iris.Context, id string) (err error)
	SearchBlog(ctx iris.Context, id string) (out []*model.Blog, err error)
	NewBlogId(ctx iris.Context) string
}
type blogService struct {
	blogRepo repo.BlogRepo
	node     *snowflake.Node
}

func NewblogService(blogRepo repo.BlogRepo, node *snowflake.Node) BlogSerivce {
	return &blogService{
		blogRepo: blogRepo,
		node:     node,
	}
}

func (j *blogService) CreateBlog(ctx iris.Context, data *model.Blog) (out *model.Blog, err error) {
	data.Id = j.NewBlogId(ctx)
	out, err = j.blogRepo.Create(ctx, data)

	if err != nil {

	}

	return out, err

}

func (j *blogService) GetBlog(ctx iris.Context) (out []*model.Blog, err error) {

	out, err = j.blogRepo.Get(ctx)

	if err != nil {

	}

	return out, err

}

func (j *blogService) UpdateBlog(ctx iris.Context) (out []*model.Blog, err error) {

	out, err = j.blogRepo.Get(ctx)

	if err != nil {

	}

	return out, err

}

func (j *blogService) DeleteBlog(ctx iris.Context, id string) (err error) {

	err = j.blogRepo.DeleteById(ctx, id)

	if err != nil {
		return err
	}

	return nil

}

func (j *blogService) SearchBlog(ctx iris.Context, id string) (out []*model.Blog, err error) {

	out, err = j.blogRepo.Get(ctx)

	if err != nil {

	}

	return out, err

}

func (j *blogService) NewBlogId(ctx iris.Context) string {

	return j.node.Generate().String()

}

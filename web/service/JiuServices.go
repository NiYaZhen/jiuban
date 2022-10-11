package service

import (
	"jiuban/model"
	"jiuban/repo"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/kataras/iris/v12"
)

type JiuSerivce interface {
	CreateJiu(ctx iris.Context, data *model.Jiu) (out *model.Jiu, err error)
	GetJiu(ctx iris.Context) (out []*model.Jiu, err error)
	UpdateJiu(ctx iris.Context, id string, in *model.Jiu) (err error)
	DeleteJiu(ctx iris.Context, id string) (err error)
	SearchJiu(ctx iris.Context, id string) (out *model.Jiu, err error)
	NewJiuId(ctx iris.Context) string
	JoinJiu(ctx iris.Context, id string, userid string) (err error)
}
type jiuService struct {
	jiuRepo repo.JiuRepo
	node    *snowflake.Node
}

func NewjiuService(jiuRepo repo.JiuRepo, node *snowflake.Node) JiuSerivce {
	return &jiuService{
		jiuRepo: jiuRepo,
		node:    node,
	}
}

func (j *jiuService) CreateJiu(ctx iris.Context, data *model.Jiu) (out *model.Jiu, err error) {
	data.Id = j.NewJiuId(ctx)
	data.CreatedAt = time.Now()
	data.UpdateAt = time.Now()
	out, err = j.jiuRepo.Create(ctx, data)

	if err != nil {

	}

	return out, err

}

func (j *jiuService) GetJiu(ctx iris.Context) (out []*model.Jiu, err error) {

	out, err = j.jiuRepo.Get(ctx)

	if err != nil {

	}

	return out, err

}

func (j *jiuService) UpdateJiu(ctx iris.Context, id string, in *model.Jiu) (err error) {
	jiu, err := j.SearchJiu(ctx, id)

	if err != nil {
		return err
	}
	if in.Title != "" {
		jiu.Title = in.Title
	}
	if in.Content != "" {
		jiu.Content = in.Content
	}

	if in.Type != "" {
		jiu.Type = in.Type
	}

	if in.Remark != "" {
		jiu.Remark = in.Remark
	}
	jiu.UpdateAt = time.Now()
	err = j.jiuRepo.Update(ctx, id, jiu)

	return nil

}

func (j *jiuService) DeleteJiu(ctx iris.Context, id string) (err error) {

	err = j.jiuRepo.DeleteById(ctx, id)

	if err != nil {
		return err
	}

	return nil

}

func (j *jiuService) SearchJiu(ctx iris.Context, id string) (out *model.Jiu, err error) {

	out, err = j.jiuRepo.SearchById(ctx, id)

	if err != nil {

	}

	return out, err

}

func (j *jiuService) NewJiuId(ctx iris.Context) string {

	return j.node.Generate().String()

}

func (j *jiuService) JoinJiu(ctx iris.Context, id string, userid string) (err error) {
	jiu, err := j.SearchJiu(ctx, id)

	joinernumber := len(jiu.JoinerList)

	if joinernumber > int(jiu.PeopleNumber) {
		return err
	} else if joinernumber < int(jiu.PeopleNumber) {
		joiner := new(model.Joiner)

		joiner.Id = userid

		jiu.JoinerList = append(jiu.JoinerList, joiner)
		jiu.UpdateAt = time.Now()
		err = j.jiuRepo.Update(ctx, id, jiu)

		return nil
	}
	return nil
}

func (j *jiuService) AgreeJoinJiu(ctx iris.Context, userid string) bool {
	//ans, _ := repo.UserRepo.GetById(ctx, userid)
	return true
}

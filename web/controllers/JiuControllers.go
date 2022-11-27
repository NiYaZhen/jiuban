package controllers

import (
	"fmt"
	"jiuban/model"
	"jiuban/repo"
	"jiuban/web/service"

	"github.com/kataras/iris/v12"
)

type JiuController interface {
	PostCreateJiu(ctx iris.Context)
	DeleteBy(ctx iris.Context, id string)
	PutBy(ctx iris.Context, id string)
	GetJiu(ctx iris.Context)
	GetBy(ctx iris.Context, id string)
	PutJoinBy(ctx iris.Context, id string, userid string)
}
type jiuController struct {
	jiuRepo    repo.JiuRepo
	jiuService service.JiuSerivce
}

func NewjiuController(jiuRepo repo.JiuRepo, jiuService service.JiuSerivce) JiuController {
	return &jiuController{
		jiuRepo:    jiuRepo,
		jiuService: jiuService,
	}
}

func (j *jiuController) PostCreateJiu(ctx iris.Context) {

	jius := new(model.Jiu)
	err := ctx.ReadJSON(&jius)

	insertResult, err := j.jiuService.CreateJiu(ctx, jius)
	fmt.Println(insertResult)

	if err != nil {
		panic(err.Error())
	}

	ctx.JSON(insertResult)

}

func (j *jiuController) DeleteBy(ctx iris.Context, id string) {
	id = ctx.Params().Get("param2")
	err := j.jiuService.DeleteJiu(ctx, id)

	if err != nil {
		panic(err.Error())
	}

}
func (j *jiuController) PutBy(ctx iris.Context, id string) {
	jius := new(model.Jiu)
	err := ctx.ReadJSON(&jius)
	fmt.Println(err)
	id = ctx.Params().Get("param2")
	fmt.Println(jius)
	err = j.jiuService.UpdateJiu(ctx, id, jius)

	if err != nil {
		panic(err.Error())
	}

}

func (j *jiuController) GetJiu(ctx iris.Context) {

	insertResult, err := j.jiuService.GetJiu(ctx)

	if err != nil {
		panic(err.Error())
	}

	ctx.JSON(insertResult)

}

func (j *jiuController) GetBy(ctx iris.Context, id string) {
	fmt.Printf("aaaa")
	id = ctx.Params().Get("param2")
	fmt.Print(id)
	insertResult, err := j.jiuService.SearchJiu(ctx, id)

	if err != nil {
		panic(err.Error())
	}

	ctx.JSON(insertResult)

}

func (j *jiuController) PutJoinBy(ctx iris.Context, id string, userid string) {
	jius := new(model.Jiu)
	err := ctx.ReadJSON(&jius)
	fmt.Println(err)
	id = ctx.Params().Get("param2")
	userid = ctx.Params().Get("param3")
	fmt.Println(jius)

	err = j.jiuService.JoinJiu(ctx, id, userid)

	if err != nil {
		panic(err.Error())
	}

}

func (j *jiuController) GetTypeJiu(ctx iris.Context) {
	searchType := new(model.SearchType)
	err := ctx.ReadJSON(&searchType)
	insertResult, err := j.jiuService.GetTypeJiu(ctx, searchType)

	if err != nil {
		panic(err.Error())
	}

	ctx.JSON(insertResult)

}

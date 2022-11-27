package service

import (
	"encoding/base64"
	"fmt"
	"jiuban/model"
	"jiuban/repo"
	"os"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/kataras/iris/v12"
)

type JiuSerivce interface {
	CreateJiu(ctx iris.Context, data *model.Jiu) (out *model.Jiu, err error)
	GetJiu(ctx iris.Context) (out []*model.Jiu, err error)
	GetTypeJiu(ctx iris.Context, searchType *model.SearchType) (out []*model.Jiu, err error)
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
	data.ImgUrl = SetImgBase64(ctx, data.ImgUrl)
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

func (j *jiuService) GetTypeJiu(ctx iris.Context, searchType *model.SearchType) (out []*model.Jiu, err error) {

	out, err = j.jiuRepo.SearchType(ctx, searchType.SearchType)

	if err != nil {
		return nil, err
	}

	return out, nil

}

func SetImgBase64(ctx iris.Context, imgList []string) []string {
	var (
		enc  = base64.StdEncoding
		path string
	)
	for i, img := range imgList {
		if img[11] == 'j' {
			img = img[23:]
			path = fmt.Sprintf("/img/%s%d.jpg", "qqq", i)
		} else if img[11] == 'p' {
			img = img[22:]
			path = fmt.Sprintf("img/%s%d.png", "qqq", i)
		} else if img[11] == 'g' {
			img = img[22:]
			path = fmt.Sprintf("img/%s%d.gif", "qqq", i)
		} else {
			fmt.Println("不支持該文件類型")
		}

		data, err := enc.DecodeString(img)
		if err != nil {
			fmt.Println(err.Error())
		}

		f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer f.Close()
		f.Write(data)
		path = "http://localhost:8080" + path
		imgList[i] = path
		fmt.Println(path)
	}

	return imgList
}

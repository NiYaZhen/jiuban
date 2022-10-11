package repo

import (
	"fmt"
	"jiuban/db"
	"jiuban/model"
	"log"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

type JiuRepo interface {
	Create(ctx iris.Context, Jiu *model.Jiu) (out *model.Jiu, err error)
	Get(ctx iris.Context) (out []*model.Jiu, err error)
	DeleteById(ctx iris.Context, id string) error
	Update(ctx iris.Context, id string, in *model.Jiu) error
	SearchById(ctx iris.Context, id string) (out *model.Jiu, err error)
}

type jiurepo struct {
}

func NewJiuRepo() JiuRepo {
	return &jiurepo{}
}

func (r *jiurepo) Create(ctx iris.Context, data *model.Jiu) (out *model.Jiu, err error) {
	out = new(model.Jiu)
	fmt.Println(data)
	_, err = db.CJiu.InsertOne(ctx, data)

	if err != nil {
		panic(err.Error())
	}

	out.Id = data.Id
	out.Title = data.Title
	out.Content = data.Content
	out.Type = data.Type
	out.Remark = data.Remark
	out.PeopleNumber = data.PeopleNumber
	out.CreatedAt = data.CreatedAt
	out.UpdateAt = data.UpdateAt

	return out, err

}

func (r *jiurepo) Get(ctx iris.Context) (out []*model.Jiu, err error) {

	cursor, err := db.CJiu.Find(ctx, bson.D{{}})

	var results []*model.Jiu
	if err := cursor.All(ctx, &results); err != nil {

	}

	return results, nil

}

func (r *jiurepo) DeleteById(ctx iris.Context, id string) error {
	filter := bson.M{"id": id}
	d, _ := db.CJiu.DeleteOne(ctx, filter)
	if d == nil {
		return nil
	}
	return nil
}

func (r *jiurepo) Update(ctx iris.Context, id string, in *model.Jiu) error {
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": in,
	}
	result, _ := db.CJiu.UpdateOne(ctx, filter, update)
	if result == nil {
		return nil
	}
	return nil
}

func (r *jiurepo) SearchById(ctx iris.Context, id string) (out *model.Jiu, err error) {
	var data *model.Jiu
	filter := bson.M{"id": id}
	d := db.CJiu.FindOne(ctx, filter)
	e := d.Decode(&data)
	if e != nil {
		log.Fatal(e)
	}

	return data, nil
}

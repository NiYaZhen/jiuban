package repo

import (
	"fmt"
	"jiuban/db"
	"jiuban/model"
	"log"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

type BlogRepo interface {
	Create(ctx iris.Context, Blog *model.Blog) (out *model.Blog, err error)
	Get(ctx iris.Context) (out []*model.Blog, err error)
	DeleteById(ctx iris.Context, id string) error
	Update(ctx iris.Context, in *model.Blog) error
	SearchById(ctx iris.Context, id string) (out []*model.Blog, err error)
}

type blogrepo struct {
}

func NewBlogRepo() BlogRepo {
	return &blogrepo{}
}

func (r *blogrepo) Create(ctx iris.Context, data *model.Blog) (out *model.Blog, err error) {
	out = new(model.Blog)
	fmt.Println(data)
	_, err = db.CBlog.InsertOne(ctx, data)

	if err != nil {
		panic(err.Error())
	}

	out.Id = data.Id
	out.Title = data.Title
	out.Content = data.Content
	out.CreatedAt = data.CreatedAt
	out.UpdateAt = data.UpdateAt

	return out, err

}

func (r *blogrepo) Get(ctx iris.Context) (out []*model.Blog, err error) {

	cursor, err := db.CBlog.Find(ctx, bson.D{{}})

	var results []*model.Blog
	if err := cursor.All(ctx, &results); err != nil {

	}

	return results, nil

}

func (r *blogrepo) DeleteById(ctx iris.Context, id string) error {
	filter := bson.M{"id": id}
	d, _ := db.CBlog.DeleteOne(ctx, filter)
	if d == nil {
		return nil
	}
	return nil
}

func (r *blogrepo) Update(ctx iris.Context, in *model.Blog) error {
	filter := bson.M{"id": in.Id}

	result, _ := db.CBlog.UpdateOne(ctx, filter, in)
	if result == nil {
		return nil
	}
	return nil
}

func (r *blogrepo) SearchById(ctx iris.Context, id string) (out []*model.Blog, err error) {
	var result []*model.Blog
	filter := bson.M{"id": id}
	d := db.CBlog.FindOne(ctx, filter)
	e := d.Decode(&result)
	if e != nil {
		log.Fatal(e)
	}

	return result, nil
}

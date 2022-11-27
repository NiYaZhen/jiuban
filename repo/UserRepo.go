package repo

import (
	"fmt"
	"jiuban/db"
	"jiuban/model"
	"time"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepo interface {
	Create(ctx iris.Context, data *model.User) (out *model.User, err error)
	Get(ctx iris.Context, email string) (*model.User, error)
	GetById(ctx iris.Context, userid string) (*model.User, error)
	Update(ctx iris.Context, user *model.User) error
	UpdateNewPassword(ctx iris.Context, email, key, hashpassword string) (*model.User, error)
	GetOtherEmail(ctx iris.Context, otheremail string) (*model.User, error)
}

type userrepo struct {
}

func NewUserRepo() UserRepo {
	return &userrepo{}
}

func (r *userrepo) Create(ctx iris.Context, data *model.User) (out *model.User, err error) {
	out = new(model.User)
	fmt.Println(data)
	_, err = db.CUser.InsertOne(ctx, data)

	if err != nil {
		panic(err.Error())
	}

	out.Id = data.Id
	out.Key = data.Key
	out.Name = data.Name
	out.Email = data.Email
	out.Password = data.Password
	out.CreatedAt = data.CreatedAt
	out.UpdatedAt = data.UpdatedAt

	return out, err

}

func (r *userrepo) Get(ctx iris.Context, email string) (*model.User, error) {
	ans := new(model.User)
	filter := bson.M{"email": email}
	if err := db.CUser.FindOne(ctx, filter).Decode(ans); err != nil {
		return ans, err

	}

	return ans, nil

}

func (r *userrepo) GetOtherEmail(ctx iris.Context, otheremail string) (*model.User, error) {
	ans := new(model.User)
	filter := bson.M{"otheremail": otheremail}
	if err := db.CUser.FindOne(ctx, filter).Decode(ans); err != nil {
		return ans, err

	}

	return ans, nil

}

func (r *userrepo) Update(ctx iris.Context, user *model.User) error {

	filter := bson.M{"email": user.Email}
	user.UpdatedAt = time.Now()
	update := bson.M{"$set": user}
	err, _ := db.CUser.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}

	return nil

}

func (r *userrepo) UpdateNewPassword(ctx iris.Context, email, key, hashpassword string) (*model.User, error) {
	var ans *model.User
	filter := bson.M{"email": email}
	if err := db.CUser.FindOne(ctx, filter).Decode(&ans); err == nil {
		fmt.Println(err)
		ans.Key = key
		ans.Password = hashpassword

	}
	fmt.Println("aaaaaaaa")
	fmt.Println(ans)
	r.Update(ctx, ans)
	return ans, nil

}

func (r *userrepo) GetById(ctx iris.Context, userid string) (*model.User, error) {
	ans := new(model.User)
	filter := bson.M{"id": userid}
	if err := db.CUser.FindOne(ctx, filter).Decode(ans); err != nil {
		return ans, err

	}

	return ans, nil

}

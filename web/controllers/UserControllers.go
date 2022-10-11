package controllers

import (
	"fmt"
	"jiuban/model"
	"jiuban/repo"
	"jiuban/web/service"

	"github.com/kataras/iris/v12"
)

type UserController interface {
	PostLogin(ctx iris.Context)
}
type userController struct {
	userRepo    repo.UserRepo
	userService service.UserSerivce
}

func NewUserController(userRepo repo.UserRepo, userService service.UserSerivce) UserController {
	return &userController{
		userRepo:    userRepo,
		userService: userService,
	}
}

func (j *userController) PostLogin(ctx iris.Context) {

	user := new(model.User)
	err := ctx.ReadJSON(&user)
	fmt.Println(err)
	if user.Email == "" && user.Password == "" {
		ctx.WriteString("輸入點東西啊？\n")
	} else if user.Email == "" {
		ctx.WriteString("沒帳號怎麼登入？\n")
	} else if user.Password == "" {
		ctx.WriteString("你忘了你的密碼了？？\n")
	} else if j.userService.VerifyEmail(user.Email) == false {
		ctx.WriteString("電子郵件格式錯誤\n")
	} else {
		j.userService.Login(ctx, user.Email, user.Password)
	}

}

func (j *userController) PostRegister(ctx iris.Context) {

	user := new(model.User)
	err := ctx.ReadJSON(&user)
	if user.Email == "" && user.Password == "" {
		ctx.WriteString("輸入點東西啊？\n")
	} else if user.Email == "" {
		ctx.WriteString("沒帳號怎麼註冊？\n")
	} else if user.Password == "" {
		ctx.WriteString("你不輸入密碼？？\n")
	} else if j.userService.VerifyEmail(user.Email) == false {
		ctx.WriteString("電子郵件格式錯誤\n")
	}
	err = j.userService.Register(ctx, user.Email, user.Password, user.Name)

	if err != nil {
		panic(err.Error())
	}

	ctx.JSON(user)

}

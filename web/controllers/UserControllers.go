package controllers

import (
	"fmt"
	"jiuban/middleware"
	"jiuban/model"
	"jiuban/repo"
	"jiuban/web/service"

	"github.com/kataras/iris/v12"
)

type UserController interface {
	PostLogin(ctx iris.Context)
	PostRegister(ctx iris.Context)
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
		ctx.JSON(&model.UserViewRes{
			Err:      1,
			Emessage: "user signin failed",
			Message:  "登入失敗，未輸入資料",
			Data:     &model.User{},
		})
	} else if user.Email == "" {
		ctx.JSON(&model.UserViewRes{
			Err:      1,
			Emessage: "user signin failed",
			Message:  "登入失敗，未輸入電子郵件",
			Data:     &model.User{},
		})
	} else if user.Password == "" {
		ctx.JSON(&model.UserViewRes{
			Err:      1,
			Emessage: "user signin failed",
			Message:  "登入失敗，未輸入密碼",
			Data:     &model.User{},
		})
	} else if j.userService.VerifyEmail(user.Email) == false {
		ctx.JSON(&model.UserViewRes{
			Err:      1,
			Emessage: "user signin failed",
			Message:  "登入失敗，電子郵件格式錯誤",
			Data:     &model.User{},
		})
	} else {
		j.userService.Login(ctx, user.Email, user.Password)
	}

}

func (j *userController) PostRegister(ctx iris.Context) {

	user := new(model.User)
	err := ctx.ReadJSON(&user)
	if user.Email == "" && user.Password == "" {
		ctx.JSON(&model.UserViewRes{
			Err:      1,
			Emessage: "user signup failed",
			Message:  "註冊失敗，未輸入資料",
			Data:     &model.User{},
		})
	} else if user.Email == "" {
		ctx.JSON(&model.UserViewRes{
			Err:      1,
			Emessage: "user signup failed",
			Message:  "註冊失敗，未輸入電子郵件",
			Data:     &model.User{},
		})
	} else if user.Password == "" {
		ctx.JSON(&model.UserViewRes{
			Err:      1,
			Emessage: "user signup failed",
			Message:  "註冊失敗，未輸入密碼",
			Data:     &model.User{},
		})
	} else if j.userService.VerifyEmail(user.Email) == false {
		ctx.JSON(&model.UserViewRes{
			Err:      1,
			Emessage: "user signup failed",
			Message:  "註冊失敗，電子郵件格式錯誤",
			Data:     &model.User{},
		})
	}
	err = j.userService.Register(ctx, user.Email, user.Password, user.Name, user.OtherEmail)

	if err != nil {
		panic(err.Error())
	}

	ctx.JSON(user)

}

func (j *userController) PostForgotPassword(ctx iris.Context) {

	user := new(model.User)
	var token string
	err := ctx.ReadJSON(&user)
	if user.OtherEmail == "" {
		ctx.JSON("沒email怎麼重置密碼？\n")
	} else {
		err = j.userService.ForgotPassword(ctx, user.OtherEmail)
		token = middleware.GetTokenHandler(ctx, user.OtherEmail)
	}

	if err != nil {
		panic(err.Error())
	}

	ctx.JSON(user)
	ctx.JSON(token)

}

func (j *userController) PutSetPassword(ctx iris.Context) {

	user := new(model.User)
	err := ctx.ReadJSON(&user)
	if user.Password == "" {
		ctx.JSON("沒新密碼怎麼重置密碼？\n")
	} else {
		email := middleware.AuthToken(ctx)
		err = j.userService.UpdateSetPassword(ctx, email, user.Password)
	}

	if err != nil {
		panic(err.Error())
	}

	ctx.JSON(user)

}

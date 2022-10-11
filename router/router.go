package router

import (
	"jiuban/repo"
	c "jiuban/web/controllers"
	"jiuban/web/service"
	"jiuban/web/snowflake"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func NewRouter(app *iris.Application) {
	Snowflake := snowflake.NewNode()
	JiuRepo := repo.NewJiuRepo()
	JiuService := service.NewjiuService(JiuRepo, Snowflake)
	UserRepo := repo.NewUserRepo()
	UserService := service.NewUserService(UserRepo, Snowflake)

	bathUrl := "/api"
	mvc.New(app.Party(bathUrl + "/jiu")).Handle(c.NewjiuController(JiuRepo, JiuService))
	mvc.New(app.Party(bathUrl + "/user")).Handle(c.NewUserController(UserRepo, UserService))

}

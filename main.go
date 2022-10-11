package main

import (
	"jiuban/db"
	"jiuban/router"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func init() {
	db.LoadTheEnv()
	db.CreateDBInstance()
}

func main() {

	app := iris.New()
	app.Use(Cors)
	app.Logger().SetLevel("debug")
	router.NewRouter(app)
	app.Use(recover.New())
	app.Use(logger.New())
	app.AllowMethods(iris.MethodOptions)

	app.Listen(":8080", iris.WithOptimizations)

}

// Cors
func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}

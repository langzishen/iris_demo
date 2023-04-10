package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/view"
	"iris-demo/application/app_session"
	"iris-demo/application/route"
)

var app *iris.Application

func main() {
	app = iris.Default()
	app.RegisterView(viewEngine())      //使用下面的方法,可以自定义模板函数  模板中使用方法:{{funcName 变量}}
	app.HandleDir("static", "./static") //定义静态文件的请求路径与实际路径的映射

	//设置前台pc home部分的路由
	new(route.Home).SetRoutes(app)

	defer iris.RegisterOnInterrupt(func() {
		app_session.Sessdb_Home.Close() //关闭session的本地数据库存储
	})
	app.Listen(":8080")
}

/**
视图环境,添加视图自定义方法
*/
func viewEngine() *view.HTMLEngine {
	HtmlEngine := iris.HTML("./application/view", ".html")

	//HtmlEngine.AddFunc("default", functions.TemplateFunc_Default) //添加视图模板自定义方法
	HtmlEngine.Reload(true)
	return HtmlEngine
}

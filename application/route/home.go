package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris_demo/application/app_session"
	"iris_demo/application/controller/home"
)

type Home struct{}

func (_ *Home) SetRoutes(app *iris.Application) {
	homePart := app.Party("/home", home.HomeAuth) //home分组,home下分组都使用homeAuth中间件,进行登录以及权限判定
	homePart.Use(iris.Compression)
	/**
	* 使用session中间件,否则在session实例:session.Start(ctx)时,报如下[debug]提示
	*[DBUG] 2023/03/19 14:20 binding: session is nil
		Maybe inside HandleHTTPError? Register it with app.UseRouter(sess.Handler()) to fix it
	*/
	homePart.UseRouter(app_session.Sess_Home.Handler()) //使用session中间件

	home_index := mvc.New(homePart.Party("/index"))
	home_index.Handle(new(home.Index))
	home_public := mvc.New(homePart.Party("/public"))
	home_public.Handle(new(home.Public))

}

package home

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/sessions"
	"iris_demo/application/app_session"
	"net/http"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var session_home = new(sessions.Session) //定义一个home的session实例变量,用来存储home路由下的session

/**
 * 登录权限验证中间件
 */
func HomeAuth(ctx iris.Context) {
	app_session.Sess_Home.UseDatabase(app_session.Sessdb_Home) //开启session的本地储存，防止长期服务时session丢失
	session_home = app_session.Sess_Home.Start(ctx, func(ctx *context.Context, c *http.Cookie, op uint8) {
		//如果config中配置了网站的域名和端口号,,则手动设置cookie的domain值,否则chrome浏览器的cookie会因为nginx的反向代理与实际访问的url地址不一致而设置cookie失败,,,iris的cookie-domain属性默认是ip:port即app.Listen的IP:port
		if host := "www.demo.com"; host != "" {
			c.Domain = host
		}
	}) //开启session
	fmt.Println("中间件运行了....")
	ctx.Next() //中间件预留,继续向下运行
	//fmt.Println("Next后......")
}

type Demo struct {
}

//实现 mvc BaseController 接口的BeginRequest和EndRequest方法
type Base struct {
	Ctx               iris.Context
	RequestApp        string            //访问的当前app
	RequestController string            //访问的当前控制器
	RequestAction     string            //访问的当前操作
	Session           *sessions.Session //home模块的session实例
	Login_id          int               //登录者Id ,即为user表id,0为未登录
	//LoginInfo         dao.Users
	Demos Demo
	//Demos2   []Demo
	LoginMap map[string]any
}

func (this *Base) BeginRequest(ctx iris.Context) {
	/**
	 ** if LoginInfo/Demos/Demos2 ... is a type of struct , error with "schema: invalid path "aaa"schema: invalid path "aaa"" in POST
	 ** But LoginMap is a type of int/string/map[string]any ... No problem
	 ** I don't know why?
	 */
	fmt.Println("BeginRequest......")
	this.Ctx = ctx
	this.Session = session_home //将session实例传给base controller基类用于控制器调用
	pathArr := strings.Split(ctx.Path(), "/")
	this.RequestApp = pathArr[1]        //app
	this.RequestController = pathArr[2] //controller
	this.RequestAction = pathArr[3]     //action
	var err error
	if this.Login_id, err = this.Session.GetInt("login_id"); err == nil && this.Login_id != -1 {
		this.initialize()
	} else {
		if this.RequestController != "public" && this.RequestAction != "login" {
			//跳转到登录
			ctx.Redirect("/" + this.RequestApp + "/public/login")
		}
	}
}

func (this *Base) EndRequest(ctx iris.Context) {}

/**
初始化
*/
func (this *Base) initialize() {
	this.Ctx.ViewData("RequestApp", this.RequestApp)
	this.Ctx.ViewData("RequestController", this.RequestController)
	this.Ctx.ViewData("RequestAction", this.RequestAction)
	this.Ctx.ViewData("login_id", this.Login_id)
	/**
	if this.Login_id > 0 {
		if this.LoginInfo.Id == 0 {
			this.LoginInfo = dao.Users{}
		}
	}
	this.Ctx.ViewData("login_info", this.LoginInfo)
	**/
	if this.Login_id > 0 {
		this.LoginMap = map[string]any{"id": 1}
	}
}

/**
 *通用的视图渲染,,默认渲染与方法名同名的控制器下的视图文件
 */
func (this *Base) GeneralView(item_template ...string) {
	fun_name := ""
	if len(item_template) == 0 {
		pc, _, _, _ := runtime.Caller(1)
		fun_name = runtime.FuncForPC(pc).Name()
		fun_name = filepath.Ext(fun_name) //去掉路径和包名部分,只保留方法名
		fun_name = strings.TrimLeft(fun_name, ".")
		if fun_name[:3] == "Get" {
			fun_name = strings.Replace(fun_name, "Get", "", 1)
		} else if fun_name[:4] == "Post" {
			fun_name = strings.Replace(fun_name, "Post", "", 1)
		}
		matchNonAlphaNumeric := regexp.MustCompile(`[^a-zA-Z0-9]+`)
		matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
		matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
		fun_name = matchNonAlphaNumeric.ReplaceAllString(fun_name, "_")  //非常规字符转化为 _
		fun_name = matchFirstCap.ReplaceAllString(fun_name, "${1}_${2}") //拆分出连续大写
		fun_name = matchAllCap.ReplaceAllString(fun_name, "${1}_${2}")   //拆分单词
		fun_name = strings.ToLower(fun_name)                             //全部转小写
		template := fun_name + ".html"
		this.Ctx.View(this.RequestApp + "/" + this.RequestController + "/" + template)
	} else {
		this.Ctx.View(this.RequestApp + "/" + this.RequestController + "/" + item_template[0])
	}
}

/**
全局获取列表页的视图
*/
func (this *Base) GetIndex() {
	this.Ctx.View(this.RequestApp + "/" + this.RequestController + "/index.html")
}

package routers

import (
	"beego_code/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//自定义方法及 RESTful 规则
	beego.Router("/", &controllers.MainController{})
	beego.Router("user/welcome", &controllers.MainController{})
	beego.Router("user/get_all_user", &controllers.MainController{}, "get:GetAllUser")
	beego.Router("user/update_user", &controllers.MainController{}, "*:UpdateUser")
	beego.Router("user/delete_user", &controllers.MainController{}, "get,post:DeleteUser")
	beego.Router("user/add_user", &controllers.MainController{}, "get:AddUser;post:UpdateUser")
	beego.Router("user/get_config_data", &controllers.MainController{}, "get:GetConfigData")

	beego.Router("game/get_game_json", &controllers.GameController{}, "get:GetGameJson")

	beego.Router("json/", &controllers.JsonTestController{}, "get:Get")

	beego.Router("login", &controllers.LoginController{}, "get:Login")

	//注解路由
	ns := beego.NewNamespace("/v2",
		beego.NSNamespace("/log",
			beego.NSInclude(
				&controllers.LogController{},
			),
		),

		// beego.NSNamespace("/user",
		//     beego.NSInclude(
		//         &controllers.UserController{},
		//     ),
		// ),
	)
	beego.AddNamespace(ns)
}
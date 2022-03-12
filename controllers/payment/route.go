package payment

import "github.com/beego/beego/v2/server/web"

func SetupRoutes() {
	web.Router("/authorize", &BaseController{}, "post:Authorize")
	web.Router("/capture", &BaseController{}, "post:Capture")
	web.Router("/refund", &BaseController{}, "post:Refund")
	web.Router("/void", &BaseController{}, "post:Void")
}

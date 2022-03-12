package merchant

import "github.com/beego/beego/v2/server/web"

func SetupRoutes() {
	web.Router("/merchant/auth", &BaseController{}, "post:Authentication")
}

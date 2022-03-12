package middleware

import "github.com/beego/beego/v2/server/web"

func Register() {
	web.InsertFilter("/*", web.BeforeRouter, merchantAuthenticated)
}

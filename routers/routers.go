// Package routers setups our REST API routes
package routers

import (
	"github.com/eldadj/dgpg/controllers/merchant"
	"github.com/eldadj/dgpg/controllers/payment"
)

func init() {
	merchant.SetupRoutes()
	payment.SetupRoutes()
	/*ns := web.NewNamespace("/api",
		merchant.Routes(), payment.Routes(),
	)
	web.AddNamespace(ns)*/
}

// Package payment REST API methods exposed by merchant endpoint

package payment

import bc "github.com/eldadj/dgpg/controllers"

type BaseController struct {
	// set in authorize call and used by other methods in the package.
	AuthorizedId string
	bc.BaseController
}

// Package controllers stores RESP API functionality
package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"github.com/eldadj/dgpg/models/authorize"
)

var ErrAuthorizeCodeNotSpecified = errors.New("authorize code not specified")

// BaseController adds functionality required/shared by all controllers
type BaseController struct {
	// set in merchant:authentication call. checked by all payment package methods
	IsAuthorized bool
	web.Controller
}

// ServeError jsonify err and return it
func (c *BaseController) ServeError(errorCode int, err error) {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Ctx.ResponseWriter.WriteHeader(errorCode)
	var content []byte
	js := map[string]interface{}{
		"error": map[string]interface{}{
			"message": err.Error(),
		}, "success": false,
	}
	content, err = json.Marshal(js)
	if err != nil {
		fmt.Fprintln(c.Ctx.ResponseWriter, err.Error())
	} else {
		fmt.Fprintln(c.Ctx.ResponseWriter, string(content))
	}
}

// ValidateAuthorizeCode validates if the id is set and exists and also not voided
func ValidateAuthorizeCode(authorizeCode string) error {
	if authorizeCode == "" {
		return ErrAuthorizeCodeNotSpecified
	}
	return authorize.ValidateAuthorizeCode(authorizeCode)
}

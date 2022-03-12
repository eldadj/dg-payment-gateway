package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/eldadj/dgpg/internal/merchant"
	"strings"
)

var merchantAuthenticated = func(ctx *context.Context) {

	var errFunc = func(message string) {
		ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx.ResponseWriter.WriteHeader(401)
		js := map[string]interface{}{
			"error": map[string]interface{}{
				"message": message,
			}, "success": false,
		}

		content, err := json.Marshal(js)
		if err != nil {
			fmt.Fprintln(ctx.ResponseWriter, err.Error())
		} else {
			fmt.Fprintln(ctx.ResponseWriter, string(content))
		}
	}

	//we ignore authencate url
	if strings.HasPrefix(ctx.Input.URL(), "/merchant/auth") {
		return
	}
	//get jwt from auth header
	header := strings.Split(ctx.Input.Header("Authorization"), " ")
	if header[0] != "Bearer" {
		errFunc("authentication header not found")
	}
	if len(header) < 2 {
		errFunc("authentication token not found")
	}
	token := header[1]
	if token == "" {
		errFunc("authentication token is invalid")
	}
	//validate the jwt token and store merchantId in context
	gCtx := ctx.Request.Context()
	//fmt.Printf("b4 ctx = %+v\n", ctx)
	if err := merchant.Validate(&gCtx, token); err != nil {
		errFunc(err.Error())
	}
	ctx.Request = ctx.Request.Clone(gCtx)
}

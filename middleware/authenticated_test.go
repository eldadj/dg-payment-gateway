package middleware

import (
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestAuthenticated(t *testing.T) {
	//enable jwt validate to use key
	//$1$CFDZmigl$9544EPgnjqtPpwDsLXP7u/
	os.Setenv("JWT_SECRET_KEY", "miglgnjqtPpwDsLXP7u/")
	type args struct {
		header map[string]string
		url    string
		status int
	}
	tests := []struct {
		name      string
		args      args
		wantValue interface{}
	}{
		{
			name:      "ignore /merchant/auth",
			wantValue: "",
			args:      args{header: map[string]string{}, url: "/merchant/auth", status: 0},
		},
		{
			name:      "no header passed",
			wantValue: "authentication header not found",
			args:      args{header: map[string]string{}, url: "/authorize", status: 401},
		},
		{
			name:      "authentication token not passed",
			wantValue: "authentication token not found",
			args:      args{header: map[string]string{"Authorization": "Bearer"}, url: "/authorize", status: 401},
		},
		{
			name:      "authentication token passed is empty",
			wantValue: "authentication token is invalid",
			args:      args{header: map[string]string{"Authorization": "Bearer  "}, url: "/authorize", status: 401},
		},
		{
			name:      "authentication token passed fails validated",
			wantValue: "invalid authentication token",
			args:      args{header: map[string]string{"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtaWQiOjJ9.2UtNBvBcZJlwatvbiuFkFwWS7ZliHcIs7_ZxMTFt9sEfail"}, url: "/authorize", status: 401},
		},
		{
			name:      "authentication token passed and validates",
			wantValue: "",
			args:      args{header: map[string]string{"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtaWQiOjJ9.2UtNBvBcZJlwatvbiuFkFwWS7ZliHcIs7_ZxMTFt9sE"}, status: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//context is beego context
			ctx := context.NewContext()
			ctx.Input.Context = ctx
			//we need a responsewriter and NewRecorder() is set for testing
			rr := httptest.NewRecorder()
			ctx.ResponseWriter = &context.Response{ResponseWriter: rr}
			uri, err := url.Parse(tt.args.url)
			if err == nil {
				ctx.Request = &http.Request{
					URL:    uri,
					Header: map[string][]string{},
				}
				for key, value := range tt.args.header {
					ctx.Request.Header.Set(key, value)
				}
				merchantAuthenticated(ctx)
				assert.Equal(t, tt.args.status, ctx.ResponseWriter.Status)
				if tt.wantValue != "" {
					assert.Contains(t, rr.Body.String(), tt.wantValue)
				}
			}
		})
	}
}

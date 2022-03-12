package merchant_test

import (
	"bytes"
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
	"github.com/eldadj/dgpg/dto/merchant/authenticate"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/internal/model"
	//_ "github.com/eldadj/dgpg/routers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func (ts *TestSuite) TestAuthentication() {
	tests := []struct {
		name       string
		wantCode   int
		wantResp   string
		args       interface{}
		setRequest func() (*http.Request, error)
	}{
		{
			name:     "invalid request data",
			wantCode: 500,
			wantResp: errors.ErrInvalidRequestData.Error(),
		},
		{
			name: "invalid username",
			args: authenticate.Request{
				Username: "invalid username",
				Password: "invalid_password",
			},
			wantCode: 500,
			wantResp: model.ErrMerchantAuthenticationFailed.Error(),
		},
		{
			name: "authenticate ok",
			args: authenticate.Request{
				Username: "m1",
				Password: "merchant1",
			},
			wantCode: 200,
			wantResp: `"token": `,
		},
	}

	for _, tt := range tests {
		t := ts.T()
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			if tt.setRequest != nil {
				r, err := tt.setRequest()
				if assert.Nil(t, err) {
					web.BeeApp.Handlers.ServeHTTP(w, r)
					assert.Equal(t, tt.wantCode, w.Code)
					assert.Contains(t, w.Body.String(), tt.wantResp)
				}
			} else {
				b := []byte("")
				if tt.args != nil {
					var err error
					b, err = json.Marshal(tt.args)
					assert.Nil(t, err)
				}
				r, _ := http.NewRequest("POST", "/merchant/auth", bytes.NewReader(b))
				web.BeeApp.Handlers.ServeHTTP(w, r)
				assert.Equal(t, tt.wantCode, w.Code)
				assert.Contains(t, w.Body.String(), tt.wantResp)
				//assert.JSONEq(t, tt.wantResp, w.Body.String())
			}
		})
	}
}

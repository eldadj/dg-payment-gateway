package payment_test

import (
	"bytes"
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/dto/payment/void"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func (s *TestSuite) TestVoid() {
	cannotVoidAuthorizeCode, _, err := s.CaptureTestCreate200USDAuthorizeCapture10N50USD()
	assert.Nil(s.T(), err)
	canVoidAuthorizeCode, _, err := s.AuthoriseTestCreateUSDAuthorize()
	assert.Nil(s.T(), err)
	tests := []struct {
		name     string
		wantCode int
		req      void.Request
		wantResp string
	}{
		{
			name:     "invalid authorize code",
			wantCode: 500,
			req: void.Request{
				AuthorizeCode: payment.AuthorizeCode{Code: "1234"},
			},
			wantResp: errors.ErrAuthorizeCodeNotFound.Error(),
		},
		{
			name:     "cannot void already captured",
			wantCode: 500,
			req:      void.Request{AuthorizeCode: payment.AuthorizeCode{Code: cannotVoidAuthorizeCode}},
			wantResp: errors.ErrAuthorizeCannotVoid.Error(),
		},
		{
			name:     "void authorize code",
			wantCode: 200,
			req:      void.Request{AuthorizeCode: payment.AuthorizeCode{Code: canVoidAuthorizeCode}},
			wantResp: "",
		},
	}

	for _, tt := range tests {
		t := s.T()
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b, err := json.Marshal(tt.req)
			assert.Nil(t, err)
			r, _ := http.NewRequest("POST", "/void", bytes.NewReader(b))
			web.BeeApp.Handlers.ServeHTTP(w, r)
			assert.Equal(t, tt.wantCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantResp)
		})
	}
}

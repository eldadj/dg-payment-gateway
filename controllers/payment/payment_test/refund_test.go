package payment_test

import (
	"bytes"
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
	"github.com/eldadj/dgpg/dto"
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/dto/payment/request"
	"github.com/eldadj/dgpg/dto/payment/response"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func (s *TestSuite) TestRefund() {
	tests := []struct {
		name     string
		wantCode int
		req      request.Request
		wantResp string
	}{
		{
			name:     "invalid amount",
			wantCode: 500,
			req: request.Request{
				AuthorizeCode: payment.AuthorizeCode{Code: "1234"},
				Amount:        0,
			},
			wantResp: errors.ErrAuthorizeCodeNotFound.Error(),
		},
		{
			name:     "invalid authorize code",
			wantCode: 500,
			req: request.Request{
				AuthorizeCode: payment.AuthorizeCode{Code: ""},
				Amount:        20,
			},
			wantResp: errors.ErrAuthorizeCodeInvalid.Error(),
		},
		{
			name:     "authorize code not found",
			wantCode: 500,
			req: request.Request{
				AuthorizeCode: payment.AuthorizeCode{Code: "1234"},
				Amount:        20,
			},
			wantResp: errors.ErrAuthorizeCodeNotFound.Error(),
		},
	}

	for _, tt := range tests {
		t := s.T()
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b, err := json.Marshal(tt.req)
			assert.Nil(t, err)
			r, _ := http.NewRequest("POST", "/refund", bytes.NewReader(b))
			web.BeeApp.Handlers.ServeHTTP(w, r)
			assert.Equal(t, tt.wantCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantResp)
		})
	}

	s.T().Run("refund 20USD", func(t *testing.T) {
		amountToRefund := 10.0
		authorizeCode, resp, err := s.CaptureTestCreate200USDAuthorizeCapture10N50USD()
		assert.Nil(t, err)
		w := httptest.NewRecorder()
		req := &request.Request{
			AuthorizeCode: payment.AuthorizeCode{Code: authorizeCode},
			Amount:        amountToRefund,
		}
		b, err := json.Marshal(req)
		assert.Nil(t, err)
		r, _ := http.NewRequest("POST", "/refund", bytes.NewReader(b))
		web.BeeApp.Handlers.ServeHTTP(w, r)
		assert.Equal(t, 200, w.Code)
		wantResp := response.Response{
			AmountCurrency: payment.AmountCurrency{
				Amount: resp.Amount + amountToRefund,
			},
			Response: dto.Response{},
		}
		var gotResp response.Response
		err = json.Unmarshal(w.Body.Bytes(), &gotResp)
		if assert.Nil(t, err) {
			assert.Equal(t, wantResp.Amount, gotResp.Amount)
		}

	})
}

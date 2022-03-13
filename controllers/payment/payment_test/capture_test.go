package payment_test

import (
	"bytes"
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/dto/payment/request"
	"github.com/eldadj/dgpg/dto/payment/response"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func (s *TestSuite) TestCapture() {
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
			wantResp: errors.ErrCaptureAmount.Error(),
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
			r, _ := http.NewRequest("POST", "/capture", bytes.NewReader(b))
			web.BeeApp.Handlers.ServeHTTP(w, r)
			assert.Equal(t, tt.wantCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantResp)
		})
	}
}

func (s *TestSuite) TestCaptureSuccess() {
	s.T().Run("capture success", func(t *testing.T) {
		req := request.Request{
			AuthorizeCode: payment.AuthorizeCode{
				//code is tied to a credit card with 1000 as amount
				Code: "40000001",
			},
			Amount: 20,
		}
		w := httptest.NewRecorder()
		b, err := json.Marshal(req)
		assert.Nil(t, err)
		r, _ := http.NewRequest("POST", "/capture", bytes.NewReader(b))
		web.BeeApp.Handlers.ServeHTTP(w, r)
		if assert.Equal(t, 200, w.Code) {
			var resp response.Response
			err = json.Unmarshal(w.Body.Bytes(), &resp)
			assert.Nil(t, err)
			assert.Equal(t, 980, resp.Amount)
		}
	})
}

func (s *TestSuite) BeforeTest(suiteName, testName string) {
	if strings.EqualFold("TestCaptureSuccess", testName) {
		s.CreateTestCreditCards()
		models.ExecDBFunc(func(tx *gorm.DB) error {
			err := tx.Exec(`
insert into authorize(authorize_id, merchant_id, credit_card_id, currency,amount,authorize_code, status) values 
(4000000,2,1000000,'USD', 1000, '40000001', 'n')
`).Error
			return err
		})
	}
}

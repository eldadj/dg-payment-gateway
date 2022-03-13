package payment_test

import (
	"bytes"
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/dto/payment/authorize"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"

	"net/http"
	"testing"
)

func (s *TestSuite) TestAuthorize() {
	validRequest := &authorize.Request{
		CreditCard: payment.CreditCard{
			OwnerName: "demo user",
			Number:    "4035 5010 0000 0008",
			ExpMonth:  10,
			ExpYear:   22,
			CVV:       "123",
		},
		AmountCurrency: payment.AmountCurrency{
			Amount:   10,
			Currency: "USD",
		},
	}
	tests := []struct {
		name       string
		wantCode   int
		wantResp   string
		args       *authorize.Request
		setRequest func() (*http.Request, error)
	}{
		{
			name:     "invalid request data",
			wantCode: 500,
			wantResp: errors.ErrInvalidRequestData.Error(),
		},
		{
			name: "invalid credit card data",
			args: &authorize.Request{
				CreditCard: payment.CreditCard{
					OwnerName: "demo user",
					Number:    "1234 5678 9012 3456",
					ExpMonth:  10,
					ExpYear:   20,
					CVV:       "123",
				},
				AmountCurrency: payment.AmountCurrency{
					Amount:   200,
					Currency: "USD",
				},
			},
			wantCode: 500,
			wantResp: errors.ErrCreditCardNoInvalid.Error(),
		},
		{
			name:     "valid request but not authenticated",
			args:     validRequest,
			wantCode: 500,
			wantResp: errors.ErrAuthorizeFailed.Error(),
		},
	}

	for _, tt := range tests {
		t := s.T()
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
				r, _ := http.NewRequest("POST", "/authorize", bytes.NewReader(b))
				web.BeeApp.Handlers.ServeHTTP(w, r)
				assert.Equal(t, tt.wantCode, w.Code)
				assert.Contains(t, w.Body.String(), tt.wantResp)
				if w.Code == 200 {
					var resp authorize.Response
					err := json.Unmarshal(w.Body.Bytes(), &resp)
					if assert.Nil(t, err) {
						assert.Equal(t, tt.args.Amount, resp.Amount)
					}
				}
			}
		})
	}

	//authenticate b4 authorize
	s.T().Run("authorize success", func(t *testing.T) {
		b, err := json.Marshal(validRequest)
		r, _ := http.NewRequest("POST", "/authorize", bytes.NewReader(b))
		assert.Nil(t, err)
		//validate token so we have a merchantId stored and update our request context
		r, err = s.ValidateMerchantUpdateRequestContext(r)
		assert.Nil(t, err)
		w := httptest.NewRecorder()
		web.BeeApp.Handlers.ServeHTTP(w, r)
		if assert.Equal(t, 200, w.Code) {
			var resp authorize.Response
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			if assert.Nil(t, err) {
				assert.Equal(t, validRequest.Amount, resp.Amount)
			}
		}
	})

	/*s.T().Run("check data", func(t *testing.T) {
		w := httptest.NewRecorder()
		b := []byte("")
		var err error
		data := authorize.Request{
			CreditCard: payment.CreditCard{
				OwnerName: "demo user",
				Number:    "TEST 5678 9012 3456",
				ExpMonth:  10,
				ExpYear:   22,
				CVV:       "123",
			},
			AmountCurrency: payment.AmountCurrency{
				Amount:   10,
				Currency: "USD",
			},
		}
		b, err = json.Marshal(data)
		assert.Nil(t, err)
		r, _ := http.NewRequest("POST", "/authorize", bytes.NewReader(b))
		web.BeeApp.Handlers.ServeHTTP(w, r)
		assert.Equal(t, 200, w.Code)
		var resp authorize.Response
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, resp.Currency, data.Currency)
		assert.Equal(t, resp.Amount, data.Amount)
		assert.NotEmpty(t, resp.AuthorizeCode)
	})*/
}

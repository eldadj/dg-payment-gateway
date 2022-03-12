package payment_test

import (
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/dto/payment/authorize"
	"github.com/eldadj/dgpg/internal/errors"

	"net/http"
	"testing"
)

func (s *TestSuite) TestAuthorize() {
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
			name: "invalid credit card data",
			args: authorize.Request{
				CreditCard: payment.CreditCard{
					OwnerName: "",
				},
				AmountCurrency: payment.AmountCurrency{},
			},
			wantCode: 500,
			wantResp: payment.ErrCreditCardInvalid.Error(),
		},
		/*{
			name: "valid data",
			args: authorize.Request{
				CreditCard: payment.CreditCard{
					OwnerName: "demo user",
					Number:    "1234 5678 9012 3456 7890",
					ExpMonth:  10,
					ExpYear:   22,
					CVV:       "123",
				},
				AmountCurrency: payment.AmountCurrency{
					Amount:   10,
					Currency: "USD",
				},
			},
			wantCode: 200,
			wantResp: "",
		},*/
	}

	for _, tt := range tests {
		t := s.T()
		t.Run(tt.name, func(t *testing.T) {
			/*w := httptest.NewRecorder()
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
			}*/
		})
	}

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

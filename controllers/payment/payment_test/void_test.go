package payment_test

import (
	"net/http"
	"testing"
)

func (s *TestSuite) TestVoid() {
	tests := []struct {
		name       string
		wantCode   int
		wantResp   string
		setRequest func() (*http.Request, error)
	}{
		{
			name:     "check ok",
			wantCode: 200,
			wantResp: `{"status": "ok"}`,
		},
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
				r, _ := http.NewRequest("POST", "/void", bytes.NewReader([]byte("")))
				web.BeeApp.Handlers.ServeHTTP(w, r)
				assert.Equal(t, tt.wantCode, w.Code)
				assert.JSONEq(t, tt.wantResp, w.Body.String())
			}*/
		})
	}
}

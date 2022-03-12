package models_test

import (
	"github.com/eldadj/dgpg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func (ts *TestSuite) TestInitDB() {
	tests := []struct {
		name       string
		wantErr    assert.ErrorAssertionFunc
		wantValue  assert.ValueAssertionFunc
		beforeTest func() error
	}{
		{
			name: "db connect",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Nil(t, err)
			},
		},
	}

	for _, tt := range tests {
		t := ts.T()
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				assert.Nil(t, tt.beforeTest())
			}
			if tt.wantErr != nil {
				tt.wantErr(t, models.InitDB())
			}
			//assert.Equal(t, tt.wantCode, w.Code)
			//assert.Contains(t, w.Body.String(), tt.wantResp)
		})
	}
}

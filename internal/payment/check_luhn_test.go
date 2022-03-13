package payment

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckLuhnTest(t *testing.T) {
	tests := []struct {
		name      string
		cardNo    string
		wantValue bool
	}{
		{
			name:      "invalid luhn card no",
			cardNo:    "1234 5678 9012 3456",
			wantValue: false,
		},
		{
			name:      "valid luhn card no",
			cardNo:    "4035 5010 0000 0008",
			wantValue: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := IsCreditCardLuhnValid(tt.cardNo)
			assert.Equal(t, tt.wantValue, valid)
		})
	}
}

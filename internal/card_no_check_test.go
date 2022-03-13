package internal

import (
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCardNoCheck(t *testing.T) {
	type args struct {
		cardNo            string
		testAgainstCardNo string
		returnErr         error
	}
	tests := []struct {
		name      string
		args      args
		wantValue interface{}
	}{
		{
			name:      "testAgainstCardNo card no",
			args:      args{cardNo: "123", testAgainstCardNo: ""},
			wantValue: errors.ErrCreditCardNoInvalid,
		},
		{
			name:      "blank card no",
			args:      args{cardNo: "", testAgainstCardNo: "1234"},
			wantValue: errors.ErrCreditCardNoInvalid,
		},
		{
			name:      "card no is ok",
			args:      args{cardNo: "1234", testAgainstCardNo: "1234"},
			wantValue: nil,
		},
		{
			name:      "authorization failure card no",
			args:      args{cardNo: "4000 0000 0000 0119", testAgainstCardNo: AuthorizeFailCard, returnErr: errors.ErrAuthorizeCreditCardFailed},
			wantValue: errors.ErrAuthorizeCreditCardFailed,
		},
		{
			name:      "capture failure card no",
			args:      args{cardNo: "4000 0000 0000 0259", testAgainstCardNo: CaptureFailCard, returnErr: errors.ErrCaptureCreditCardFailed},
			wantValue: errors.ErrCaptureCreditCardFailed,
		},
		{
			name:      "refund failure card no",
			args:      args{cardNo: "4000 0000 0000 3238", testAgainstCardNo: RefundFailCard, returnErr: errors.ErrRefundCreditCardFailed},
			wantValue: errors.ErrRefundCreditCardFailed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CardNoCheck(tt.args.cardNo, tt.args.testAgainstCardNo, tt.args.returnErr)
			if tt.wantValue == nil {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, tt.wantValue.(error).Error())
			}
		})
	}
}

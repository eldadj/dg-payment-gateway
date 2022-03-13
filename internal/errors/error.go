package errors

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// LogError creates a means to log an actual error, and return another error to the caller.
// This helps to prevent leaking actual application error.
// If interfaces are used, then it makes it possible to act based on the error type
// realErr error raised by go, this is logged and client doesn't see it.
// responseErr error to return to caller
//TODO: refactor to use interfaces
func LogError(realErr, responseErr error) error {
	if realErr != nil {
		log.Error(realErr)
	}
	return responseErr
}

var (
	ErrInvalidRequestData                    = errors.New("invalid request data")
	ErrMerchantLoad                          = errors.New("error loading merchant info")
	ErrMerchantTokenExpired                  = errors.New("your login session has expired")
	ErrMerchantInvalidToken                  = errors.New("invalid authentication token")
	ErrInvalidMerchantAuthorizeCode          = errors.New("authorize code is invalid for merchant")
	ErrAuthorizeInsufficientCreditCardAmount = errors.New("insufficient amount in credit card")
	ErrAuthorizeFailed                       = errors.New("error completing authorize")
	ErrAuthorizeCreditCardFailed             = errors.New("credit card authorisation failure")
	ErrAuthorizeCodeInvalid                  = errors.New("authorize code is invalid")
	ErrFieldValueRequired                    = func(fieldName string) error {
		return errors.New(fmt.Sprintf("%s is required", fieldName))
	}
	ErrAuthorizeCodeNotFound      = errors.New("authorize code not found")
	ErrAuthorizeCodeInvalidStatus = func(status string) error {
		return errors.New(fmt.Sprintf("invalid authorize code status: %s", status))
	}
	ErrAuthorizeInvalidFieldValue = func(fieldName string) error {
		return errors.New(fmt.Sprintf("invalid value for %s in authorize ", fieldName))
	}
	ErrMerchantAuthenticationFailed = errors.New("invalid merchant username/password")

	ErrAuthorizeCannotVoid                 = errors.New("authorize code cannot be voided")
	ErrAuthorizeCannotCapture              = errors.New("authorize code is not valid for capture")
	ErrAuthorizeCannotRefund               = errors.New("authorize code is not valid for refund")
	ErrAuthorizeVoidFailed                 = errors.New("error voiding authorize code")
	ErrAuthorizeLoadCreditCard             = errors.New("error loading credit card info for authorize code")
	ErrAuthorizeAmountsAuthorize           = errors.New("error reading amounts already authorized")
	ErrCreditCardLoadAmountCurrency        = errors.New("error getting credit card amount/currency")
	ErrCreditCardUpdateAmount              = errors.New("error updating credit card amount")
	ErrCreditCardLoad                      = errors.New("error loading credit card")
	ErrCreditCardSave                      = errors.New("error loading saving card")
	ErrCreditCardNoInvalid                 = errors.New("error validating credit card number")
	ErrCapture                             = errors.New("error saving capture")
	ErrCaptureAmountExceedsAuthorizeAmount = errors.New("capture amount exceeds amount authorized")
	ErrCaptureAmount                       = errors.New("capture amount is invalid")
	ErrCaptureCreditCardFailed             = errors.New("credit card capture failure")
	ErrRefund                              = errors.New("error during refund")
	ErrRefundAmount                        = errors.New("amount cannot be refunded")
	ErrRefundCreditCardFailed              = errors.New("credit card refund failure")
)

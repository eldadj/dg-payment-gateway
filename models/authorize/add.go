package authorize

import (
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/internal/errors"
	"gorm.io/gorm"
)

// Add creates a new record in authorize table. returns new authorizeCode if created
func Add(tx *gorm.DB, merchantId int64, creditCardId int64, payAC payment.AmountCurrency) (
	authorizeCode string, err error) {
	if merchantId <= 0 {
		return "", errors.ErrAuthorizeInvalidFieldValue("merchant_id")
	}
	if creditCardId <= 0 {
		return "", errors.ErrAuthorizeInvalidFieldValue("credit_card_id")
	}
	if payAC.Currency == "" {
		return "", errors.ErrAuthorizeInvalidFieldValue("currency")
	}
	if payAC.Amount <= 0 {
		return "", errors.ErrAuthorizeInvalidFieldValue("amount")
	}
	//TODO validate if merchantId and creditCardId are valid
	a := Authorize{MerchantId: merchantId, CreditCardId: creditCardId, Currency: payAC.Currency, Amount: payAC.Amount}
	if err = tx.Create(&a).Error; err == nil {
		authorizeCode = a.AuthorizeCode
	} else {
		err = errors.LogError(err, errors.ErrAuthorizeFailed)
	}

	return
}

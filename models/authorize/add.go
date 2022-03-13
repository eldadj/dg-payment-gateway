package authorize

import (
	auth "github.com/eldadj/dgpg/dto/payment/authorize"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models/credit_card"
	"gorm.io/gorm"
)

// Add creates a new record in authorize table. returns new authorizeCode if created
func Add(tx *gorm.DB, merchantId int64, req auth.Request) (
	authorizeCode string, err error) {
	if merchantId <= 0 {
		return "", errors.ErrAuthorizeInvalidFieldValue("merchant_id")
	}
	if req.Currency == "" {
		return "", errors.ErrAuthorizeInvalidFieldValue("currency")
	}
	if req.Amount <= 0 {
		return "", errors.ErrAuthorizeInvalidFieldValue("amount")
	}
	creditCardId, err := credit_card.Add(tx, req.CreditCard, req.AmountCurrency)
	//TODO validate if merchantId and creditCardId are valid
	a := Authorize{MerchantId: merchantId, CreditCardId: creditCardId, Currency: req.Currency, Amount: req.Amount}
	if err = tx.Create(&a).Error; err == nil {
		authorizeCode = a.AuthorizeCode
	} else {
		err = errors.LogError(err, errors.ErrAuthorizeFailed)
	}

	return
}

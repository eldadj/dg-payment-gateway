package authorize

import (
	"github.com/eldadj/dgpg/dto"
	"github.com/eldadj/dgpg/dto/payment"
)

type Response struct {
	payment.AuthorizeCode
	payment.AmountCurrency
	dto.Response
}

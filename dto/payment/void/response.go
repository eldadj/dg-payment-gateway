package void

import (
	"github.com/eldadj/dgpg/dto"
	"github.com/eldadj/dgpg/dto/payment"
)

type Response struct {
	payment.AmountCurrency
	dto.Response
}

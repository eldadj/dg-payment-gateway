package credit_card

type creditCard struct {
	CreditCardId  int64 `gorm:"primaryKey"`
	OwnerName     string
	CardNo        string
	ExpMonth      int
	ExpYear       int
	CVV           string `gorm:"column:cvv"`
	CurrencyCode  string
	CurrentAmount float64
}

func (cc *creditCard) TableName() string {
	return "credit_card"
}

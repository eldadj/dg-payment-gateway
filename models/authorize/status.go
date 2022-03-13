package authorize

import (
	"strings"
)

//n = new just created,
//v = voided,
//r = refunded,
//p = capturing or when at least when capture has taken place
//c = fully captured

//CanVoid only a new authorize
func (a *Authorize) CanVoid() bool {
	return strings.EqualFold("n", a.Status) //|| strings.EqualFold("p", a.Status)
}

//CanCapture authorize that is ok/new and has not refund
func (a *Authorize) CanCapture() bool {
	//when voided
	return (strings.EqualFold("n", a.Status) || strings.EqualFold("p", a.Status)) && !a.HasRefund
}

//CanRefund only refund when not voided
func (a *Authorize) CanRefund() bool {
	return !strings.EqualFold("v", a.Status) //|| !strings.EqualFold("r", a.Status)
}

package payment

import (
	"encoding/json"
	"github.com/eldadj/dgpg/dto/payment/request"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models/refund"
)

func (c *BaseController) Refund() {
	req, err := validateRefund(c)
	if err != nil {
		c.ServeError(500, err)
		return
	}

	resp, err := refund.DoRefund(req)
	if err != nil {
		c.ServeError(500, err)
		return
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

func validateRefund(c *BaseController) (req request.Request, err error) {
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		return req, errors.LogError(err, errors.ErrInvalidRequestData)
	}
	if req.Amount < 0 {
		return req, errors.ErrCaptureAmount
	}

	return req, req.Validate()
}

package payment

import (
	"encoding/json"
	"github.com/eldadj/dgpg/dto/payment/authorize"
	"github.com/eldadj/dgpg/internal/errors"
	authorize2 "github.com/eldadj/dgpg/models/authorize"
)

func (c *BaseController) Authorize() {
	req, err := validate(c)
	if err != nil {
		c.ServeError(500, err)
		return
	}

	resp, err := authorize2.DoAuthorize(c.Ctx.Request.Context(), *req)
	if err != nil {
		c.ServeError(500, err)
		return
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// validate data is valid
func validate(c *BaseController) (*authorize.Request, error) {
	var req *authorize.Request
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		return nil, errors.LogError(err, errors.ErrInvalidRequestData)
	}
	if err := req.CreditCard.Validate(); err != nil {
		return nil, err
	}
	if err := req.AmountCurrency.Validate(); err != nil {
		return nil, err
	}

	return req, nil
}

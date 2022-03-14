package payment

import (
	"encoding/json"
	"github.com/eldadj/dgpg/dto/payment/request"
	"github.com/eldadj/dgpg/internal/errors"
	capture2 "github.com/eldadj/dgpg/models/capture"
)

func (c *BaseController) Capture() {
	req, err := validateCapture(c)
	if err != nil {
		c.ServeError(500, err)
		return
	}

	//set request merchant id
	err = SetMerchantId(c.Ctx.Request.Context(), &req.Request)
	if err != nil {
		c.ServeError(500, err)
		return
	}

	resp, err := capture2.DoCapture(req)
	if err != nil {
		c.ServeError(500, err)
		return
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

func validateCapture(c *BaseController) (req request.Request, err error) {
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		return req, errors.LogError(err, errors.ErrInvalidRequestData)
	}
	if req.Amount <= 0 {
		return req, errors.ErrCaptureAmount
	}

	return req, req.Validate()
}

package payment

import (
	"encoding/json"
	"github.com/eldadj/dgpg/dto/payment/void"
	"github.com/eldadj/dgpg/internal/errors"
	void2 "github.com/eldadj/dgpg/models/void"
)

func (c *BaseController) Void() {
	req, err := validateVoid(c)
	if err != nil {
		c.ServeError(500, err)
		return
	}
	resp, err := void2.Void(*req)
	if err != nil {
		c.ServeError(500, err)
		return
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

func validateVoid(c *BaseController) (*void.Request, error) {
	var req *void.Request
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		return nil, errors.LogError(err, errors.ErrInvalidRequestData)
	}
	return req, req.Validate()
}

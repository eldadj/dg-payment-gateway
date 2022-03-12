package merchant

import (
	"encoding/json"
	"github.com/eldadj/dgpg/dto"
	"github.com/eldadj/dgpg/dto/merchant/authenticate"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models"
)

func (c *BaseController) Authentication() {
	req, err := c.validate()
	if err != nil {
		c.ServeError(500, err)
		return
	}
	token, err := models.Authenticate(*req)
	if err != nil {
		c.ServeError(500, err)
		return
	}
	resp := authenticate.Response{Response: dto.Response{Success: true}, Token: token}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *BaseController) validate() (*authenticate.Request, error) {
	var req authenticate.Request
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		return nil, errors.LogError(err, errors.ErrInvalidRequestData)
	}
	return &req, nil
}

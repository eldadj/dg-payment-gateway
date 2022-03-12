package authenticate

import "github.com/eldadj/dgpg/dto"

type Response struct {
	dto.Response
	Token string `json:"token"`
}

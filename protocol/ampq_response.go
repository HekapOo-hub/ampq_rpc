package protocol

import (
	"encoding/json"
)

type GetResponse struct {
	User      User   `json:"user"`
	ErrorBody string `json:"error_body"`
}

func (g GetResponse) Encode() ([]byte, error) {
	return json.Marshal(g)
}

type CreateResponse struct {
	ErrorBody string `json:"error_body"`
}

func (c CreateResponse) Encode() ([]byte, error) {
	return json.Marshal(c)
}

type ResendResponse struct {
	ErrorBody string `json:"error_body"`
}

func (r ResendResponse) Encode() ([]byte, error) {
	return json.Marshal(r)
}

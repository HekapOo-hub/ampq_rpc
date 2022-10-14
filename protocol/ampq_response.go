package protocol

import (
	"ampq_rpc/internal/model"
	"encoding/json"
)

type GetResponse struct {
	User      model.User
	ErrorBody string
}

func (g GetResponse) Encode() ([]byte, error) {
	return json.Marshal(g)
}

type CreateResponse struct {
	ErrorBody string
}

func (c CreateResponse) Encode() ([]byte, error) {
	return json.Marshal(c)
}

type ResendResponse struct {
	ErrorBody string
}

func (r ResendResponse) Encode() ([]byte, error) {
	return json.Marshal(r)
}
package protocol

import (
	"encoding/json"
	"fmt"
	"regexp"
)

var (
	getRequestRegexp    = regexp.MustCompile(`{"key":".*?"}`)
	createRequestRegexp = regexp.MustCompile(`{"user":{"id":".*?","name":".*?","age":.*?}}`)
	resendRequestRegexp = regexp.MustCompile(`{"start_id":".*?","end_id":".*?","key":".*?"}`)
)

type Encoder interface {
	Encode() ([]byte, error)
}

type GetRequest struct {
	Key string `json:"key"`
}

func (g GetRequest) Encode() ([]byte, error) {
	return json.Marshal(g)
}

func IsGetRequest(data []byte) bool {
	return getRequestRegexp.Match(data)
}

func ParseGetRequest(data []byte) (GetRequest, error) {
	gr := GetRequest{}
	err := json.Unmarshal(data, &gr)
	if err != nil {
		return GetRequest{}, fmt.Errorf("parse get request: %w", err)
	}
	return gr, err
}

type CreateRequest struct {
	User User `json:"user"`
}

func (c CreateRequest) Encode() ([]byte, error) {
	return json.Marshal(c)
}

func IsCreateRequest(data []byte) bool {
	return createRequestRegexp.Match(data)
}

func ParseCreateRequest(data []byte) (CreateRequest, error) {
	cr := CreateRequest{}
	err := json.Unmarshal(data, &cr)
	if err != nil {
		return CreateRequest{}, fmt.Errorf("parse create request: %w", err)
	}
	return cr, nil
}

type ResendRequest struct {
	StartId string `json:"start_id"`
	EndId   string `json:"end_id"`

	Key string `json:"key"`
}

func (r ResendRequest) Encode() ([]byte, error) {
	return json.Marshal(r)
}

func IsResendRequest(data []byte) bool {
	return resendRequestRegexp.Match(data)
}

func ParseResendRequest(data []byte) (ResendRequest, error) {
	rr := ResendRequest{}
	err := json.Unmarshal(data, &rr)
	if err != nil {
		return ResendRequest{}, fmt.Errorf("parse resend request: %w", err)
	}
	return rr, nil
}

package protocol

import (
	"encoding/json"
	"fmt"
	"regexp"
)

var (
	getRequestRegexp    = regexp.MustCompile(`{"Key":".*?"}`)
	createRequestRegexp = regexp.MustCompile(`{"User":{"ID":".*?","Name":".*?","Age":.*?}}`)
	resendRequestRegexp = regexp.MustCompile(`{"StartId":".*?","EndId":".*?","Key":".*?"}`)
)

type Encoder interface {
	Encode() ([]byte, error)
}

type GetRequest struct {
	Key string
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
	User User
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
	StartId string
	EndId   string

	Key string
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

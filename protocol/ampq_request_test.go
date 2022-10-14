package protocol

import (
	"ampq_rpc/internal/model"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestIsCreateRequestEmpty(t *testing.T) {
	cr := CreateRequest{}
	sb := strings.Builder{}
	err := json.NewEncoder(&sb).Encode(cr)
	require.NoError(t, err)
	require.True(t, IsCreateRequest([]byte(sb.String())))
}

func TestIsCreateRequestInitialized(t *testing.T) {
	cr := CreateRequest{
		User: model.User{
			Name: "Andrew",
			Age:  123,
			ID:   uuid.New().String(),
		},
	}
	sb := strings.Builder{}
	err := json.NewEncoder(&sb).Encode(cr)
	require.NoError(t, err)
	require.True(t, IsCreateRequest([]byte(sb.String())))
}

func TestIsCreateRequestIncorrectCompare(t *testing.T) {
	gr := GetRequest{}
	sb := strings.Builder{}
	err := json.NewEncoder(&sb).Encode(gr)
	require.NoError(t, err)
	require.False(t, IsCreateRequest([]byte(sb.String())))
}

func TestIsGetRequestEmpty(t *testing.T) {
	gr := GetRequest{}
	sb := strings.Builder{}
	err := json.NewEncoder(&sb).Encode(gr)
	require.NoError(t, err)
	require.True(t, IsGetRequest([]byte(sb.String())))
}

func TestIsGetRequestInitialized(t *testing.T) {
	gr := GetRequest{
		Key: uuid.New().String(),
	}
	sb := strings.Builder{}
	err := json.NewEncoder(&sb).Encode(gr)
	require.NoError(t, err)
	require.True(t, IsGetRequest([]byte(sb.String())))
}

func TestIsGetRequestIncorrectCompare(t *testing.T) {
	cr := CreateRequest{User: model.User{Name: "Me"}}
	sb := strings.Builder{}
	err := json.NewEncoder(&sb).Encode(cr)
	require.NoError(t, err)
	require.False(t, IsGetRequest([]byte(sb.String())))
}

func TestIsResendRequestEmpty(t *testing.T) {
	rr := ResendRequest{}
	sb := strings.Builder{}
	err := json.NewEncoder(&sb).Encode(rr)
	require.NoError(t, err)
	require.True(t, IsResendRequest([]byte(sb.String())))
}

func TestIsResendRequestInitialized(t *testing.T) {
	rr := ResendRequest{
		StartId: uuid.New().String(),
		EndId:   uuid.New().String(),
		Key:     "k",
	}
	sb := strings.Builder{}
	err := json.NewEncoder(&sb).Encode(rr)
	require.NoError(t, err)
	require.True(t, IsResendRequest([]byte(sb.String())))
}

func TestIsResendRequestIncorrectCompare(t *testing.T) {
	gr := GetRequest{}
	sb := strings.Builder{}
	err := json.NewEncoder(&sb).Encode(gr)
	require.NoError(t, err)
	require.False(t, IsResendRequest([]byte(sb.String())))
}

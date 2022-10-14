package protocol

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsCreateRequestEmpty(t *testing.T) {
	cr := CreateRequest{}
	data, err := cr.Encode()
	require.NoError(t, err)
	require.True(t, IsCreateRequest(data))
}

func TestIsCreateRequestInitialized(t *testing.T) {
	cr := CreateRequest{
		User: User{
			Name: "Andrew",
			Age:  123,
			ID:   uuid.New().String(),
		},
	}
	data, err := cr.Encode()
	require.NoError(t, err)
	require.True(t, IsCreateRequest(data))
}

func TestIsCreateRequestIncorrectCompare(t *testing.T) {
	gr := GetRequest{}
	data, err := gr.Encode()
	require.NoError(t, err)
	require.False(t, IsCreateRequest(data))
}

func TestIsGetRequestEmpty(t *testing.T) {
	gr := GetRequest{}
	data, err := gr.Encode()
	require.NoError(t, err)
	require.True(t, IsGetRequest(data))
}

func TestIsGetRequestInitialized(t *testing.T) {
	gr := GetRequest{
		Key: uuid.New().String(),
	}
	data, err := gr.Encode()
	require.NoError(t, err)
	require.True(t, IsGetRequest(data))
}

func TestIsGetRequestIncorrectCompare(t *testing.T) {
	cr := CreateRequest{User: User{Name: "Me"}}
	data, err := cr.Encode()
	require.NoError(t, err)
	require.False(t, IsGetRequest(data))
}

func TestIsResendRequestEmpty(t *testing.T) {
	rr := ResendRequest{}
	data, err := rr.Encode()
	require.NoError(t, err)
	require.True(t, IsResendRequest(data))
}

func TestIsResendRequestInitialized(t *testing.T) {
	rr := ResendRequest{
		StartId: uuid.New().String(),
		EndId:   uuid.New().String(),
		Key:     "k",
	}
	data, err := rr.Encode()
	require.NoError(t, err)
	require.True(t, IsResendRequest(data))
}

func TestIsResendRequestIncorrectCompare(t *testing.T) {
	gr := GetRequest{}
	data, err := gr.Encode()
	require.NoError(t, err)
	require.False(t, IsResendRequest(data))
}

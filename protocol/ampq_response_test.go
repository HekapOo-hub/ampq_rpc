package protocol

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateResponseIsNotErrorTrue(t *testing.T) {
	r := CreateResponse{ErrorBody: ""}
	data, err := r.Encode()
	require.NoError(t, err)
	logrus.Info(string(data))
	require.True(t, IsNotError(data))
}

func TestCreateResponseIsNotErrorFalse(t *testing.T) {
	r := CreateResponse{ErrorBody: "ERROR!!!"}
	data, err := r.Encode()
	require.NoError(t, err)
	require.False(t, IsNotError(data))
}

func TestGetResponseIsNotErrorTrue(t *testing.T) {
	r := GetResponse{ErrorBody: "", User: User{ID: uuid.New().String()}}
	data, err := r.Encode()
	require.NoError(t, err)
	require.True(t, IsNotError(data))
}

func TestGetResponseIsNotErrorFalse(t *testing.T) {
	r := GetResponse{ErrorBody: "ERROR!!!!!!!", User: User{ID: uuid.New().String()}}
	data, err := r.Encode()
	require.NoError(t, err)
	require.False(t, IsNotError(data))
}

func TestResendResponseIsNotErrorTrue(t *testing.T) {
	r := ResendResponse{ErrorBody: ""}
	data, err := r.Encode()
	require.NoError(t, err)
	require.True(t, IsNotError(data))
}

func TestResendResponseIsNotErrorFalse(t *testing.T) {
	r := ResendResponse{ErrorBody: "."}
	data, err := r.Encode()
	require.NoError(t, err)
	require.False(t, IsNotError(data))
}

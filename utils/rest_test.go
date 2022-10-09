package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallWithJsonNormal(t *testing.T) {
	// testing part
	resp, err := RESTCallWithJson("v1/chain/get_info", POST, nil)

	// assertion part
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

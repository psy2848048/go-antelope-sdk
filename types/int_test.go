package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalInt(t *testing.T) {
	newInt := NewIntFromInt64(100)

	byteInt, err := json.Marshal(newInt)

	assert.NoError(t, err)
	assert.Equal(t, []byte(`"100"`), byteInt)
}

func TestUnmarshalInt(t *testing.T) {
	intByte := []byte(`"100"`)
	newInt := &Int{}

	err := json.Unmarshal(intByte, newInt)

	assert.NoError(t, err)
	assert.Equal(t, NewIntFromInt64(100), newInt)
}

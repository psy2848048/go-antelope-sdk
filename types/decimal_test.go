package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecimalMarshal(t *testing.T) {
	coin, err := NewDecimalFromInt64(1_000100, 6)
	if err != nil {
		panic(err)
	}

	byteDecimal, err := json.Marshal(coin)
	assert.NoError(t, err)
	assert.Equal(t, []byte(`"1.000100"`), byteDecimal)
}

func TestDecimalUnmarshal(t *testing.T) {
	marshalledDec := []byte(`"1.000100"`)
	dec := &Decimal{}

	err := json.Unmarshal(marshalledDec, dec)
	assert.NoError(t, err)

	expectedCoin, err := NewDecimalFromInt64(1_000100, 6)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, expectedCoin, dec)
}

func TestNewDecimalWithZeroDecimal(t *testing.T) {
	marshalledDecimal := []byte(`"1000100"`)
	dec := &Decimal{}

	err := json.Unmarshal(marshalledDecimal, dec)
	assert.NoError(t, err)

	expectedCoin, err := NewDecimalFromInt64(1_000100, 0)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, expectedCoin, dec)
	assert.Equal(t, "1000100", dec.String())
}

func TestNewDecimalWithBelowDecimal(t *testing.T) {
	marshalledDecimal := []byte(`"0.000100"`)
	dec := &Decimal{}

	err := json.Unmarshal(marshalledDecimal, dec)
	assert.NoError(t, err)

	expectedCoin, err := NewDecimalFromInt64(100, 6)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, expectedCoin, dec)
	assert.Equal(t, "0.000100", dec.String())
}

func TestNewDecimalWithBelowDecimal2(t *testing.T) {
	marshalledDec := []byte(`"0.100000"`)
	dec := &Decimal{}

	err := json.Unmarshal(marshalledDec, dec)
	assert.NoError(t, err)

	expectedDec, err := NewDecimalFromInt64(100000, 6)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, expectedDec, dec)
	assert.Equal(t, "0.100000", dec.String())
}

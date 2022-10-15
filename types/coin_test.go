package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinMarshal(t *testing.T) {
	coin, err := NewCoinFromInt64(1_000100, 6, "EOS")
	if err != nil {
		panic(err)
	}

	byteCoin, err := json.Marshal(coin)
	assert.NoError(t, err)
	assert.Equal(t, []byte(`"1.000100 EOS"`), byteCoin)
}

func TestCoinUnmarshal(t *testing.T) {
	marshalledCoin := []byte(`"1.000100 EOS"`)
	coin := &Coin{}

	err := json.Unmarshal(marshalledCoin, coin)
	assert.NoError(t, err)

	expectedCoin, err := NewCoinFromInt64(1_000100, 6, "EOS")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, expectedCoin, coin)
}

func TestNewCoinWithLowerDenom(t *testing.T) {
	// expected error in validation
	_, err := NewCoinFromInt64(1_000100, 6, "Eos")
	assert.Error(t, err)
}

func TestNewCoinWithNoDecimal(t *testing.T) {
	marshalledCoin := []byte(`"1000100 EOS"`)
	coin := &Coin{}

	err := json.Unmarshal(marshalledCoin, coin)
	assert.NoError(t, err)

	expectedCoin, err := NewCoinFromInt64(1_000100, 0, "EOS")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, expectedCoin, coin)
	assert.Equal(t, "1000100 EOS", coin.String())
}

func TestNewCoinWithBelowDecimal(t *testing.T) {
	marshalledCoin := []byte(`"0.000100 EOS"`)
	coin := &Coin{}

	err := json.Unmarshal(marshalledCoin, coin)
	assert.NoError(t, err)

	expectedCoin, err := NewCoinFromInt64(100, 6, "EOS")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, expectedCoin, coin)
	assert.Equal(t, "0.000100 EOS", coin.String())
}

func TestNewCoinWithBelowDecimal2(t *testing.T) {
	marshalledCoin := []byte(`"0.100000 EOS"`)
	coin := &Coin{}

	err := json.Unmarshal(marshalledCoin, coin)
	assert.NoError(t, err)

	expectedCoin, err := NewCoinFromInt64(100000, 6, "EOS")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, expectedCoin, coin)
	assert.Equal(t, "0.100000 EOS", coin.String())
}

package types

import (
	"encoding"
	"encoding/json"
	"fmt"
	"math/big"
)

const maxBitLen = 256

type Int struct {
	i *big.Int
}

func NewIntFromString(amt string) (*Int, error) {
	ret := &Int{}
	if intAmt, ok := new(big.Int).SetString(amt, 0); ok {
		ret.i = intAmt
		return ret, nil
	} else {
		return nil, fmt.Errorf("cannot parse into big/Int Input : %s", amt)
	}
}

func NewIntFromInt64(amt int64) *Int {
	newInt := &Int{}
	newInt.i = big.NewInt(amt)

	return newInt
}

func (i Int) MarshalJSON() ([]byte, error) {
	if i.i == nil { // Necessary since default Uint initialization has i.i as nil
		i.i = new(big.Int)
	}
	return marshalJSON(i.i)
}

// UnmarshalJSON defines custom decoding scheme
func (i *Int) UnmarshalJSON(bz []byte) error {
	if i.i == nil { // Necessary since default Int initialization has i.i as nil
		i.i = new(big.Int)
	}
	return unmarshalJSON(i.i, bz)
}

func (i Int) String() string {
	return i.i.String()
}

// MarshalJSON for custom encoding scheme
// Must be encoded as a string for JSON precision
func marshalJSON(i encoding.TextMarshaler) ([]byte, error) {
	text, err := i.MarshalText()
	if err != nil {
		return nil, err
	}

	return json.Marshal(string(text))
}

// UnmarshalJSON for custom decoding scheme
// Must be encoded as a string for JSON precision
func unmarshalJSON(i *big.Int, bz []byte) error {
	var text string
	if err := json.Unmarshal(bz, &text); err != nil {
		return err
	}

	return unmarshalText(i, text)
}

func unmarshalText(i *big.Int, text string) error {
	if err := i.UnmarshalText([]byte(text)); err != nil {
		return err
	}

	if i.BitLen() > maxBitLen {
		return fmt.Errorf("integer out of range: %s", text)
	}

	return nil
}

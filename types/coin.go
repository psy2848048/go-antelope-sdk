package types

import (
	"bytes"
	"fmt"
	"strings"
)

type Coin struct {
	Amount  Int
	Decimal uint
	Denom   string
}

func NewCoinFromInt64(amt int64, decimal uint, denom string) (*Coin, error) {
	newCoin := &Coin{
		Amount:  *NewIntFromInt64(amt),
		Decimal: decimal,
		Denom:   denom,
	}

	err := newCoin.Validate()
	if err != nil {
		return nil, err
	}

	return newCoin, nil
}

func (c Coin) String() string {
	strAmt := c.Amount.String()

	// decimal point
	var strAmtWithDecimal string
	if len(strAmt) > int(c.Decimal) && c.Decimal > 0 {
		strAmtWithDecimal = fmt.Sprintf("%s.%s", strAmt[:len(strAmt)-int(c.Decimal)], strAmt[len(strAmt)-int(c.Decimal):])
	} else if len(strAmt) <= int(c.Decimal) && c.Decimal > 0 {
		zero := strings.Repeat("0", int(c.Decimal)-len(strAmt))
		strAmtWithDecimal = fmt.Sprintf("0.%s%s", zero, strAmt)
	} else {
		strAmtWithDecimal = strAmt
	}

	ret := fmt.Sprintf("%s %s", strAmtWithDecimal, c.Denom)

	return ret
}

func (c Coin) MarshalJSON() ([]byte, error) {
	strCoin := c.String()

	return []byte(fmt.Sprintf(`"%s"`, strCoin)), nil
}

// UnmarshalJSON defines custom decoding scheme
func (c *Coin) UnmarshalJSON(bz []byte) error {
	// split amount & denom
	splitedBytes := bytes.Split(bz, []byte(" "))

	interimAmt := strings.ReplaceAll(string(splitedBytes[0]), `"`, "")
	c.Denom = strings.ReplaceAll(string(splitedBytes[1]), `"`, "")

	decimalParts := strings.Split(interimAmt, ".")
	if len(decimalParts) == 1 {
		c.Decimal = 0

		newInt, err := NewIntFromString(decimalParts[0])
		if err != nil {
			return err
		}

		c.Amount = *newInt
	} else {
		c.Decimal = uint(len(decimalParts[1]))

		strAmt := strings.Join(decimalParts, "")

		breakPoint := 0
		for idx, val := range strAmt {
			if fmt.Sprintf("%c", val) == "0" && idx == 0 {
				breakPoint = idx
			} else if fmt.Sprintf("%c", val) != "0" && idx == 0 {
				breakPoint = -1
				break
			} else if fmt.Sprintf("%c", val) == "0" && idx > 0 {
				breakPoint = idx
			} else if fmt.Sprintf("%c", val) != "0" && idx > 0 {
				break
			}
		}

		if breakPoint >= 0 {
			strAmt = strAmt[breakPoint+1:]
		}

		newInt, err := NewIntFromString(strAmt)
		if err != nil {
			return err
		}

		c.Amount = *newInt
	}

	err := c.Validate()

	return err
}

func (c Coin) Validate() error {
	if c.Denom != strings.ToUpper(c.Denom) {
		return fmt.Errorf("demon should be upper case, received %s", c.Denom)
	}

	return nil
}

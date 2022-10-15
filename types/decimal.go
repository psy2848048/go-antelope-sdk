package types

import (
	"fmt"
	"strings"
)

type Decimal struct {
	Amount  Int
	Decimal uint
}

func NewDecimalFromInt64(amt int64, decimal uint) (*Decimal, error) {
	newDecimal := &Decimal{
		Amount:  *NewIntFromInt64(amt),
		Decimal: decimal,
	}

	return newDecimal, nil
}

func (d Decimal) String() string {
	strAmt := d.Amount.String()

	// decimal point
	var strDecimal string
	if len(strAmt) > int(d.Decimal) && d.Decimal > 0 {
		strDecimal = fmt.Sprintf("%s.%s", strAmt[:len(strAmt)-int(d.Decimal)], strAmt[len(strAmt)-int(d.Decimal):])
	} else if len(strAmt) <= int(d.Decimal) && d.Decimal > 0 {
		zero := strings.Repeat("0", int(d.Decimal)-len(strAmt))
		strDecimal = fmt.Sprintf("0.%s%s", zero, strAmt)
	} else {
		strDecimal = strAmt
	}

	return strDecimal
}

func (d Decimal) MarshalJSON() ([]byte, error) {
	strDecimal := d.String()

	return []byte(fmt.Sprintf(`"%s"`, strDecimal)), nil
}

// UnmarshalJSON defines custom decoding scheme
func (d *Decimal) UnmarshalJSON(bz []byte) error {
	in := string(bz)
	in = strings.ReplaceAll(in, `"`, "")

	decimalParts := strings.Split(in, ".")
	if len(decimalParts) == 1 {
		d.Decimal = 0

		newInt, err := NewIntFromString(decimalParts[0])
		if err != nil {
			return err
		}

		d.Amount = *newInt
	} else {
		d.Decimal = uint(len(decimalParts[1]))

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

		d.Amount = *newInt
	}

	return nil
}

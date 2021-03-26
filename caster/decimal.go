package caster

import (
	"math"
)

func DecodeDecimal(buffer []byte, offset int, optional bool) (float64, int, bool) {
	exponent, offset, nullable := DecodeInt(buffer, offset, optional)

	// only optional
	if nullable {
		return 0, offset, true
	}

	mantissa, offset, _ := DecodeInt(buffer, offset, false)
	value := float64(mantissa) * math.Pow10(exponent)

	return float64(value), offset, false
}

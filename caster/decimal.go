package caster

import "math"

func DecodeDecimal(buffer []byte, offset int, optional bool) (float64, int, bool) {
	exponent, offset, _ := DecodeInteger(buffer, offset, optional)

	// only optional
	if optional && exponent > 0 {
		exponent--
	}

	mantissa, offset, _ := DecodeInteger(buffer, offset, false)
	value := float64(mantissa) * math.Pow10(exponent)

	return float64(value), offset, false
}

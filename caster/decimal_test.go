package caster

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

// 1. Decimal Example – Mandatory Positive Decimal
// <decimal id="1" presence="mandatory" name="Value"/>
func TestDecodeDecimalSample1(t *testing.T) {
	buffer := []byte{0b10000010, 0b00111001, 0b01000101, 0b10100011}
	value, offset, nullable := DecodeDecimal(buffer, 0, false)

	Equal(t, 94275500.0, value, "value")
	Equal(t, 4, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// 2. Decimal Example – Mandatory Positive Decimal with Scaled Mantissa
// <decimal id="1" presence="mandatory" name="Value"/>
func TestDecodeDecimalSample2(t *testing.T) {
	buffer := []byte{0b10000001, 0b00000100, 0b00111111, 0b00110100, 0b11011110}
	value, offset, nullable := DecodeDecimal(buffer, 0, false)

	Equal(t, 94275500.0, value, "value")
	Equal(t, 5, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// 3. Decimal Example – Optional Positive Decimal
// <decimal id="1" presence="optional" name="Value"/>
func TestDecodeDecimalSample3(t *testing.T) {
	buffer := []byte{0b10000011, 0b00111001, 0b01000101, 0b10100011}
	value, offset, nullable := DecodeDecimal(buffer, 0, true)

	Equal(t, float64(9427550), value, "value")
	Equal(t, 4, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// 4. Decimal Example – Mandatory Positive Decimal
// <decimal id="1" presence="mandatory" name="Value"/>
func TestDecodeDecimalSample4(t *testing.T) {
	buffer := []byte{0b11111110, 0b00111001, 0b01000101, 0b10100011}
	value, offset, nullable := DecodeDecimal(buffer, 0, false)

	Equal(t, float64(9427.550000000001), value, "value")
	Equal(t, 4, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// 5. Decimal Example – Optional Negative Decimal
// <decimal id="1" presence="optional" name="Value"/>
func TestDecodeDecimalSample5(t *testing.T) {
	buffer := []byte{0b11111110, 0b01000110, 0b00111010, 0b11011101}
	value, offset, nullable := DecodeDecimal(buffer, 0, true)

	Equal(t, float64(-9427.550000000001), value, "value")
	Equal(t, 4, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// 8. Decimal Example – Optional Negative Decimal with sign bit extension
// <decimal id="1" presence="optional" name="Value"/>
func TestDecodeDecimalSample6(t *testing.T) {
	buffer := []byte{0b11111101, 0b01111111, 0b00111111, 0b11111111}
	value, offset, nullable := DecodeDecimal(buffer, 0, true)

	Equal(t, -8.193, value, "value")
	Equal(t, 4, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

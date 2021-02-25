package caster

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

// 1. uInt32 Example – Optional Number
// <uInt32 id="1" presence="optional" name="Value"/>

func TestDecodeUIntSample1(t *testing.T) {
	buffer := []byte{0b10000000}
	value, offset, nullable := DecodeUInt(buffer, 0, true)

	Equal(t, 0, value, "value")
	Equal(t, 1, offset, "offset")
	Equal(t, true, nullable, "nullable")
}

func TestDecodeUIntSample2(t *testing.T) {
	buffer := []byte{0b10000001}
	value, offset, nullable := DecodeUInt(buffer, 0, true)

	Equal(t, 0, value, "value")
	Equal(t, 1, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

func TestDecodeUIntSample3(t *testing.T) {
	buffer := []byte{0b10000010}
	value, offset, nullable := DecodeUInt(buffer, 0, true)

	Equal(t, 1, value, "value")
	Equal(t, 1, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

func TestDecodeUIntSample4(t *testing.T) {
	buffer := []byte{0b00111001, 0b01000101, 0b10100100}
	value, offset, nullable := DecodeUInt(buffer, 0, true)

	Equal(t, 942755, value, "value")
	Equal(t, 3, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// 2. uInt32 Example – Mandatory Number
// <uInt32 id="1" presence="mandatory" name="Value"/>

func TestDecodeUIntSample5(t *testing.T) {
	buffer := []byte{0b10000000}
	value, offset, nullable := DecodeUInt(buffer, 0, false)

	Equal(t, 0, value, "value")
	Equal(t, 1, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

func TestDecodeUIntSample6(t *testing.T) {
	buffer := []byte{0b10000001}
	value, offset, nullable := DecodeUInt(buffer, 0, false)

	Equal(t, 1, value, "value")
	Equal(t, 1, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

func TestDecodeUIntSample7(t *testing.T) {
	buffer := []byte{0b00111001, 0b01000101, 0b10100011}
	value, offset, nullable := DecodeUInt(buffer, 0, false)

	Equal(t, 942755, value, "value")
	Equal(t, 3, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// Custom samples

func TestDecodeUIntSample8(t *testing.T) {
	buffer := []byte{0b01010000, 0b11000010}
	value, offset, nullable := DecodeUInt(buffer, 0, false)

	Equal(t, 10306, value, "value")
	Equal(t, 2, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

func TestDecodeUIntSample9(t *testing.T) {
	buffer := []byte{0b01010000, 0b11000010}
	value, offset, nullable := DecodeUInt(buffer, 0, true)

	Equal(t, 10305, value, "value")
	Equal(t, 2, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

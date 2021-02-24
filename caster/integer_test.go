package caster

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

// 1. Int32 Example – Optional Positive Number
// <int32 id="1" presence="optional" name="Value"/>
func TestDecodeIntegerSample1(t *testing.T) {
	buffer := []byte{0b00111001, 0b01000101, 0b10100100}
	value, offset, nullable := DecodeInteger(buffer, 0, true)

	Equal(t, 942755, value, "value")
	Equal(t, 3, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// 2. Int32 Example – Mandatory Positive Number
// <int32 id="1" presence="mandatory" name="Value"/>
func TestDecodeIntegerSample2(t *testing.T) {
	buffer := []byte{0b00111001, 0b01000101, 0b10100011}
	value, offset, nullable := DecodeInteger(buffer, 0, false)

	Equal(t, 942755, value, "value")
	Equal(t, 3, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// 3. Int32 Example – Optional Negative Number
// <int32 id="1" presence="optional" name="Value"/>
func TestDecodeIntegerSample3(t *testing.T) {
	buffer := []byte{0b001000110, 0b00111010, 0b11011101}
	value, offset, nullable := DecodeInteger(buffer, 0, true)

	Equal(t, -942755, value, "value")
	Equal(t, 3, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// 4. Int32 Example – Mandatory Negative Number
// <int32 id="1" presence="mandatory" name="Value"/>
func TestDecodeIntegerSample4(t *testing.T) {
	buffer := []byte{0b01111100, 0b00011011, 0b00011011, 0b10011101}
	value, offset, nullable := DecodeInteger(buffer, 0, false)

	Equal(t, -7942755, value, "value")
	Equal(t, 4, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// 5. Int32 Example – Mandatory Positive Number with sign-bit extension
// <int32 id="1" presence="mandatory" name="Value"/>
func TestDecodeIntegerSample5(t *testing.T) {
	buffer := []byte{0b00000000, 0b01000000, 0b10000001}
	value, offset, nullable := DecodeInteger(buffer, 0, false)

	Equal(t, 8193, value, "value")
	Equal(t, 3, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// 6. Int32 Example – Mandatory Negative Number with sign-bit extension
// <int32 id="1" presence="mandatory" name="Value"/>
func TestDecodeIntegerSample6(t *testing.T) {
	buffer := []byte{0b01111111, 0b00111111, 0b11111111}
	value, offset, nullable := DecodeInteger(buffer, 0, false)

	Equal(t, -8193, value, "value")
	Equal(t, 3, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// 1. uInt32 Example – Optional Number
// <uInt32 id="1" presence="optional" name="Value"/>

func TestDecodeIntegerSample7(t *testing.T) {
	buffer := []byte{0b10000000}
	value, offset, nullable := DecodeInteger(buffer, 0, true)

	Equal(t, 0, value, "value")
	Equal(t, 1, offset, "offset")
	Equal(t, true, nullable, "nullable")
}

func TestDecodeIntegerSample8(t *testing.T) {
	buffer := []byte{0b10000001}
	value, offset, nullable := DecodeInteger(buffer, 0, true)

	Equal(t, 0, value, "value")
	Equal(t, 1, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

func TestDecodeIntegerSample9(t *testing.T) {
	buffer := []byte{0b10000010}
	value, offset, nullable := DecodeInteger(buffer, 0, true)

	Equal(t, 1, value, "value")
	Equal(t, 1, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

func TestDecodeIntegerSample10(t *testing.T) {
	buffer := []byte{0b00111001, 0b01000101, 0b10100100}
	value, offset, nullable := DecodeInteger(buffer, 0, true)

	Equal(t, 942755, value, "value")
	Equal(t, 3, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// 2. uInt32 Example – Mandatory Number
// <uInt32 id="1" presence="mandatory" name="Value"/>

func TestDecodeIntegerSample11(t *testing.T) {
	buffer := []byte{0b10000000}
	value, offset, nullable := DecodeInteger(buffer, 0, false)

	Equal(t, 0, value, "value")
	Equal(t, 1, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

func TestDecodeIntegerSample12(t *testing.T) {
	buffer := []byte{0b10000001}
	value, offset, nullable := DecodeInteger(buffer, 0, false)

	Equal(t, 1, value, "value")
	Equal(t, 1, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

func TestDecodeIntegerSample13(t *testing.T) {
	buffer := []byte{0b00111001, 0b01000101, 0b10100011}
	value, offset, nullable := DecodeInteger(buffer, 0, false)

	Equal(t, 942755, value, "value")
	Equal(t, 3, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

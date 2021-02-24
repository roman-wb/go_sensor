package caster

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

// 1. US ASCII string Example – Optional String
// <string id="1" presence="optional" name="Value"/>

func TestDecodeStringSample1(t *testing.T) {
	buffer := []byte{0b10000000}
	value, offset, nullable := DecodeString(buffer, 0, true)

	Equal(t, "", value, "value")
	Equal(t, 1, offset, "offset")
	Equal(t, true, nullable, "nullable")
}

func TestDecodeStringSample2(t *testing.T) {
	buffer := []byte{0b01000001, 0b01000010, 0b11000011}
	value, offset, nullable := DecodeString(buffer, 0, true)

	Equal(t, "ABC", value, "value")
	Equal(t, 3, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

func TestDecodeStringSample3(t *testing.T) {
	buffer := []byte{0b00000000, 0b10000000}
	value, offset, nullable := DecodeString(buffer, 0, true)

	Equal(t, "", value, "value")
	Equal(t, 2, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// 2. US ASCII string Example – Mandatory String
// <string id="1" presence="mandatory" name="Value"/>

func TestDecodeStringSample4(t *testing.T) {
	buffer := []byte{0b01000001, 0b01000010, 0b11000011}
	value, offset, nullable := DecodeString(buffer, 0, false)

	Equal(t, "ABC", value, "value")
	Equal(t, 3, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

func TestDecodeStringSample5(t *testing.T) {
	buffer := []byte{0b10000000}
	value, offset, nullable := DecodeString(buffer, 0, false)

	Equal(t, "", value, "value")
	Equal(t, 1, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

package caster

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

// 1. byteVector Example – Optional byteVector
// <byteVector id="1" presence="optional" name="Value"/>

func TestByteVectorSample1(t *testing.T) {
	buffer := []byte{0b10000000}
	value, offset, nullable := DecodeByteVector(buffer, 0, true)

	Equal(t, []byte{}, value, "value")
	Equal(t, 1, offset, "offset")
	Equal(t, true, nullable, "nullable")
}

func TestByteVectorSample2(t *testing.T) {
	buffer := []byte{0b10000100, 0b01000001, 0b01000010, 0b01000011}
	value, offset, nullable := DecodeByteVector(buffer, 0, true)

	Equal(t, []byte("ABC"), value, "value")
	Equal(t, 4, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

func TestByteVectorSample3(t *testing.T) {
	buffer := []byte{0b10000001}
	value, offset, nullable := DecodeByteVector(buffer, 0, true)

	Equal(t, []byte{}, value, "value")
	Equal(t, 1, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

// 2. byteVector Example – Mandatory byteVector
// <byteVector id="1" presence="mandatory" name="Value"/>

func TestByteVectorSample4(t *testing.T) {
	buffer := []byte{0b10000011, 0b01000001, 0b01000010, 0b01000011}
	value, offset, nullable := DecodeByteVector(buffer, 0, false)

	Equal(t, []byte("ABC"), value, "value")
	Equal(t, 4, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

func TestByteVectorSample5(t *testing.T) {
	buffer := []byte{0b10000000}
	value, offset, nullable := DecodeByteVector(buffer, 0, false)

	Equal(t, []byte{}, value, "value")
	Equal(t, 1, offset, "offset")
	Equal(t, false, nullable, "nullable")
}

package caster

func DecodeByteVector(buffer []byte, offset int, optional bool) ([]byte, int, bool) {
	if buffer[offset] == STOP_BIT_POSITION {
		offset++
		return []byte{}, offset, optional
	}

	length, offset, _ := DecodeInt(buffer, offset, optional)
	value := buffer[offset : length+1]

	return value, offset + length, false
}

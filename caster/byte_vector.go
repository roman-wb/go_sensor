package caster

func DecodeByteVector(buffer []byte, offset int, optional bool) (string, int, bool) {
	if buffer[offset] == STOP_BIT_POSITION {
		offset++
		return "", offset, optional
	}

	length, offset, _ := DecodeInteger(buffer, offset, optional)
	value := buffer[offset : length+1]

	return string(value), offset + length, false
}

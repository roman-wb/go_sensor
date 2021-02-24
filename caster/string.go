package caster

func DecodeString(buffer []byte, offset int, optional bool) (string, int, bool) {
	if buffer[offset] == STOP_BIT_POSITION {
		offset++
		return "", offset, optional
	}

	if buffer[offset] == 0b00000000 && buffer[offset+1] == STOP_BIT_POSITION {
		offset += 2
		return "", offset, false
	}

	start := offset
	for (buffer[offset] & STOP_BIT_POSITION) == 0 {
		offset++
	}
	buffer[offset] &= STOP_BIT_RESET
	offset++

	return string(buffer[start:offset]), offset, false
}

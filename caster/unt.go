package caster

func DecodeUInt(buffer []byte, offset int, optional bool) (int, int, bool) {
	value := 0

	if buffer[offset] == STOP_BIT_POSITION {
		offset++
		return value, offset, optional
	}

	for (buffer[offset] & STOP_BIT_POSITION) == 0 {
		value = (value << DATA_BIT_SHIFT) | int(buffer[offset])
		offset++
	}
	value = (value << DATA_BIT_SHIFT) | int(buffer[offset]&STOP_BIT_RESET)
	offset++

	// only optional
	if optional && value > 0 {
		value -= 1
	}

	return value, offset, false
}

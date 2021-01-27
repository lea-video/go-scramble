package scrambler

import "math"

var feedbackMasks = [...]uint32 {
	// 2 to 8
	0x3, 0x6, 0x9, 0x1D, 0x36, 0x69, 0xA6,
	// 9 - 16, and so on
	0x17C, 0x32D, 0x4F2, 0xD34, 0x1349, 0x2532, 0x6699, 0xD295,
	0x12933, 0x2C93E, 0x593CA, 0xAFF95, 0x12B6BC, 0x2E652E, 0x5373D6, 0x9CCDAE,
	0x12BA74D, 0x36CD5A7, 0x4E5D793, 0xF5CDE95, 0x1A4E6FF2, 0x29D1E9EB, 0x7A5BC2E3, 0xB4BCD35C,
}
func getMask(highestBitPos uint8) uint32 {
	if highestBitPos < 2 || highestBitPos > 32 {
		return 0;
	}
	return feedbackMasks[highestBitPos - 2]
}

func bitCount(maximum uint32) uint8 {
	var r uint8 = 0
	for maximum != 0 {
		maximum >>= 1
		r++
	}
	return r
}

func mod(a int, b uint32) uint32 {
	return uint32(a % int(b))
}

func minOne(a uint32) uint32 {
	var n uint32 = 1
	if a > n {
		n = a
	}
	return n
}

func min(a, b uint32) uint32 {
	var k uint32 = math.MaxUint32
	if a < k {
		k = a
	}
	if b < k {
		k = b
	}
	return k
}

func calcOffset(outOffset int, max uint32) uint32 {
	m := int(max)
	offset := outOffset % m
	if offset < 0 {
		// wrap negative to positive
		offset += m
	}
	return uint32(offset)
}

// TODO: check that this does never return zeros
func offset(cur, offset, maximum uint32) uint32 {
	k := cur + offset
	if k > maximum {
		k -= maximum
	}
	return k
}

func shiftForward(cur, mask uint32) uint32 {
	if cur & 1 == 1 {
		return (cur >> 1) ^ mask
	} else {
		return cur >> 1
	}
}

func shiftBackward(cur, feedbackMask, bitMask uint32) uint32 {
	if cur & bitMask != 0 {
		return ((feedbackMask ^ cur) << 1) + 1
	} else {
		return cur << 1
	}
}
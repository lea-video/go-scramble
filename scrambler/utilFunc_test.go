package scrambler

import (
	"strconv"
	"testing"
)

func TestScrambler_getMask(t *testing.T) {
	// out of range returns 0
	gm := getMask(0)
	if gm != 0 {
		t.Error("Expected 0 when passing 0");
	}
	gm = getMask(1)
	if gm != 0 {
		t.Error("Expected 0 when passing 1");
	}
	gm = getMask(2)
	if gm == 0 {
		t.Error("Expected non-0 when passing 2");
	}
	gm = getMask(32)
	if gm == 0 {
		t.Error("Expected non-0 when passing 32");
	}
	gm = getMask(33)
	if gm != 0 {
		t.Error("Expected 0 when passing 33");
	}
	// test 2 random returns
	gm = getMask(12)
	if gm != 0xD34 {
		t.Error("Expected 0xD34 when passing 12");
	}
	gm = getMask(20)
	if gm != 0xAFF95 {
		t.Error("Expected 0xAFF95 when passing 20");
	}
}

func TestScramble_bitCount(t *testing.T) {
	n, _ := strconv.ParseUint("1011", 2, 32)
	b := bitCount(uint32(n))
	if b != 4 {
		t.Error("Expected 1011 to be a 4 bit number")
	}
	n, _ = strconv.ParseUint("0011", 2, 32)
	b = bitCount(uint32(n))
	if b != 2 {
		t.Error("Expected 0011 to be a 2 bit number")
	}
	n, _ = strconv.ParseUint("110011", 2, 32)
	b = bitCount(uint32(n))
	if b != 6 {
		t.Error("Expected 110011 to be a 6 bit number")
	}
}

func TestScrambler_min(t *testing.T) {
	// TODO: implement
}

func TestScrambler_wrappingMod(t *testing.T) {
	// TODO: implement
}

func TestScrambler_offset(t *testing.T) {
	// Check that offset never return zeros
	var i uint32 = 0
	for i < 16 {
		if offset(i,4,8) == 0 {
			t.Error("offset should not return 0")
		}
		i++
	}
	// check that it offsets correctly
	if offset(4, 2, 8) != 6 {
		t.Error("Expected 4 + 2 = 6")
	}
	// check that it wrapps correct
	if offset(4, 6, 8) != 2 {
		t.Error("Expected offset 4 by 6 to be 2")
	}
	if offset(6, 4, 8) != 2 {
		t.Error("Expected offset 6 by 4 to be 2")
	}
}

func TestScrambler_shiftForward(t *testing.T) {
	// TODO: implement
}

func TestScrambler_shiftBackward(t *testing.T) {
	// TODO: implement
}

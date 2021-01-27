package scrambler

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

const MaxInt = math.MaxUint32

func HashCode(str string) int {
	if len(str) == 0 {
		return rand.Int()
	}

	chars := []rune(str)
	hash := 0
	i := 0
	for i < len(str) {
		char := chars[i]
		hash = ((hash << 5) - hash) + int(char)
		hash |= 0;
		i++
	}
	return hash
}

func New(opts ConstructorOpts) (*Scrambler, error) {
	if opts.Maximum < 2 || opts.Maximum > MaxInt {
		return nil, errors.New(fmt.Sprintf("\"maximum\" must be a number between 2 and %d inclusive", MaxInt))
	}
	// analyse the maximum bits
	bc := bitCount(opts.Maximum)
	// calculate the 0-position from seed
	// build scrambler
	p := Scrambler{
		opts: opts,
		start: wrappingMod(opts.Seed, opts.Maximum) + 1,
		highestBitMask: 2 << (bc - 2),
		feedbackMask: getMask(bc),
		outOffset: wrappingMod(opts.OutOffset, opts.Maximum),
		atStart: true,
		atEnd: false,
	}
	p.Reset()
	return &p, nil
}

func (s *Scrambler) Reset()  {
	s.cur = s.start
	s.n = -1
}

func (s *Scrambler) All() []uint32 {
	s.Reset()
	r := make([]uint32, s.opts.Maximum)

	var i uint32 = 0
	for i < s.opts.Maximum {
		r[i] = s.Next()
		i++
	}

	return r
}

func (s *Scrambler) Next() uint32 {
	if s.atEnd {
		return 0
	}
	s.atStart = false

	s.cur = shiftForward(s.cur, s.feedbackMask)
	for s.cur > s.opts.Maximum {
		s.cur = shiftForward(s.cur, s.feedbackMask)
	}

	if s.cur == s.start && !s.opts.Loop {
		s.atEnd = true
	}

	s.n++
	if s.n > int(s.opts.Maximum - 1) {
		s.n -= int(s.opts.Maximum)
	}

	return offset(s.cur, s.outOffset, s.opts.Maximum)
}

func (s *Scrambler) Prev() uint32 {
	if s.atStart {
		return 0
	}
	s.atEnd = false

	s.cur = shiftBackward(s.cur, s.feedbackMask, s.highestBitMask)
	for s.cur > s.opts.Maximum {
		s.cur = shiftBackward(s.cur, s.feedbackMask, s.highestBitMask)
	}

	if s.cur == s.start && !s.opts.Loop {
		s.atStart = true
		return 0
	}

	s.n--
	if s.n < 0 {
		s.n += int(s.opts.Maximum)
	}

	return offset(s.cur, s.outOffset, s.opts.Maximum)
}

func (s *Scrambler) Current() uint32 {
	if s.n == -1 {
		return 0
	} else {
		return offset(s.cur, s.outOffset, s.opts.Maximum)
	}
}

func (s *Scrambler) CurrentPos() int {
	return s.n
}

func (s *Scrambler) Lookahead(n uint32) []uint32 {
	var out []uint32
	if s.opts.Loop {
		out = make([]uint32, n)
	} else if s.n == -1 {
		out = make([]uint32, min(n, s.opts.Maximum))
	} else {
		out = make([]uint32, min(n, s.opts.Maximum - uint32(s.n) - 1))
	}

	c := s.Clone()
	i := 0
	for i < len(out) {
		out[i] = c.Next()
		i++
	}

	return out
}

func (s *Scrambler) Lookbehind(n uint32) []uint32 {
	var out []uint32
	if s.n == -1 {
		return make([]uint32, 0)
	}

	if s.opts.Loop {
		out = make([]uint32, n)
	} else {
		out = make([]uint32, min(n, uint32(s.n)))
	}

	c := s.Clone()
	i := 0
	for i < len(out) {
		out[i] = c.Prev()
		i++
	}

	return out
}

func (s *Scrambler) Clone() *Scrambler {
	c, e := New(s.opts)
	if e != nil {
		panic(e)
	}
	c.goTo(s.cur, s.n)
	return c
}

func (s *Scrambler) goTo(cur uint32, n int)  {
	s.cur = cur
	s.n = n
}
